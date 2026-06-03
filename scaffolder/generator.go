package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run generator.go <project-name> <type: clean|hex|vertical|modular|standard> [with-ui]")
		os.Exit(1)
	}

	projectName := os.Args[1]
	archType := strings.ToLower(os.Args[2])
	withUI := false
	if len(os.Args) > 3 && os.Args[3] == "with-ui" {
		withUI = true
	}

	basePath := "/output/" + projectName

	fmt.Printf("Generating project '%s' with architecture '%s' (UI: %v)...\n", projectName, archType, withUI)

	// Base directories
	var dirs []string
	switch archType {
	case "clean":
		dirs = []string{"cmd/app", "internal/entity", "internal/usecase/interfaces", "internal/infrastructure/db", "internal/infrastructure/config", "internal/delivery/http", "pkg", "api", "configs"}
	case "hex":
		dirs = []string{"cmd/app", "internal/core/domain", "internal/core/ports/input", "internal/core/ports/output", "internal/service", "internal/adapters/input/http", "internal/adapters/output/persistence", "api", "configs"}
	case "vertical":
		dirs = []string{"cmd/server", "internal/features", "internal/shared/database", "internal/shared/middleware", "internal/platform", "api"}
	case "standard":
		dirs = []string{"cmd/app", "pkg", "internal", "api", "configs"}
	case "modular":
		dirs = []string{"cmd/app", "internal/modules/user", "internal/modules/auth", "internal/shared/db", "internal/shared/logger", "api", "configs"}
	default:
		fmt.Printf("Invalid architecture type: %s\n", archType)
		os.Exit(1)
	}

	// Add frontend directories if requested
	if withUI {
		dirs = append(dirs, "web/src/assets", "web/src/components", "web/src/features", "web/public")
	} else {
		dirs = append(dirs, "web") // Empty web dir for consistency
	}

	for _, dir := range dirs {
		path := filepath.Join(basePath, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	// --- GENERATE FILES ---

	// 1. main.go with optional Embed
	mainPath := filepath.Join(basePath, "cmd", "app", "main.go")
	if archType == "vertical" {
		mainPath = filepath.Join(basePath, "cmd", "server", "main.go")
	}

	mainImports := `"fmt"
	"os"
	"os/signal"
	"syscall"
	"gopkg.in/yaml.v3"`
	
	embedLogic := ""
	if withUI {
		mainImports += "\n\t\"embed\"\n\t\"net/http\"\n\t\"io/fs\""
		embedLogic = `
//go:embed web/dist/*
var frontendContent embed.FS

func serveFrontend() {
	distFS, _ := fs.Sub(frontendContent, "web/dist")
	http.Handle("/", http.FileServer(http.FS(distFS)))
	fmt.Println("UI served at http://localhost:8080")
}`
	}

	mainContent := fmt.Sprintf(`package main

import (
	%s
)

type Config struct {
	Server struct {
		Port int    ` + "`" + `yaml:"port"` + "`" + `
		Host string ` + "`" + `yaml:"host"` + "`" + `
	} ` + "`" + `yaml:"server"` + "`" + `
}

%s

func main() {
	conf := &Config{}
	data, err := os.ReadFile("configs/config.yaml")
	if err == nil {
		yaml.Unmarshal(data, conf)
	}

	fmt.Printf("Project initialized on port %%d\n", conf.Server.Port)

	%s

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	fmt.Println("\nShutting down server gracefully...")
}
`, mainImports, embedLogic, strings.TrimSpace(func() string { if withUI { return "go serveFrontend()" }; return "" }()))

	if err := os.WriteFile(mainPath, []byte(mainContent), 0644); err != nil {
		fmt.Printf("Error writing main.go: %v\n", err)
		os.Exit(1)
	}

	// 2. React Boilerplate (if with-ui)
	if withUI {
		pkgJson := `{
  "name": "` + projectName + `-ui",
  "version": "0.1.0",
  "type": "module",
  "scripts": {
    "dev": "vite",
    "build": "tsc && vite build",
    "lint": "eslint . --ext ts,tsx --report-unused-disable-directives --max-warnings 0",
    "preview": "vite preview"
  }
}`
		if err := os.WriteFile(filepath.Join(basePath, "web/package.json"), []byte(pkgJson), 0644); err != nil {
			fmt.Printf("Error writing web/package.json: %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile(filepath.Join(basePath, "web/src/App.tsx"), []byte(`export default function App() { return <h1>`+projectName+` UI</h1> }`), 0644); err != nil {
			fmt.Printf("Error writing web/src/App.tsx: %v\n", err)
			os.Exit(1)
		}
	}

	// 3. Makefile with Frontend support
	uiBuildStep := ""
	if withUI {
		uiBuildStep = "build-ui: ## Build the React frontend\n\tcd web && npm install && npm run build\n\n"
	}

	makefileContent := `BINARY_NAME=` + projectName + `

.PHONY: help build run test lint clean docker-build build-ui

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

` + uiBuildStep + `build: ` + func() string { if withUI { return "build-ui " }; return "" }() + `## Build the application binary
	go build -o bin/$(BINARY_NAME) ` + func() string { if archType == "vertical" { return "cmd/server/main.go" }; return "cmd/app/main.go" }() + `

run: ## Run the application
	go run ` + func() string { if archType == "vertical" { return "cmd/server/main.go" }; return "cmd/app/main.go" }() + `

test: ## Run unit tests
	go test ./...

lint: ## Run linter
	golangci-lint run

docker-build: ## Build production Docker image
	docker build -t $(BINARY_NAME) .
`
	if err := os.WriteFile(filepath.Join(basePath, "Makefile"), []byte(makefileContent), 0644); err != nil {
		fmt.Printf("Error writing Makefile: %v\n", err)
		os.Exit(1)
	}

	// 4. Standard Configs
	if err := os.WriteFile(filepath.Join(basePath, "configs/config.yaml"), []byte("server:\n  port: 8080\n"), 0644); err != nil {
		fmt.Printf("Error writing configs/config.yaml: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(filepath.Join(basePath, "go.mod"), []byte("module github.com/organization/"+projectName+"\n\ngo 1.24\n"), 0644); err != nil {
		fmt.Printf("Error writing go.mod: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(filepath.Join(basePath, ".gitignore"), []byte("bin/\nnode_modules/\ndist/\n"), 0644); err != nil {
		fmt.Printf("Error writing .gitignore: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Project generated successfully in /output/" + projectName)
}
