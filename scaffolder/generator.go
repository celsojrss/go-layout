package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run generator.go <project-name> <type: clean|hex|vertical|modular|standard>")
		os.Exit(1)
	}

	projectName := os.Args[1]
	archType := strings.ToLower(os.Args[2])
	basePath := "/output/" + projectName

	fmt.Printf("Generating project '%s' with architecture '%s'...\n", projectName, archType)

	var dirs []string
	switch archType {
	case "clean":
		dirs = []string{
			"cmd/app",
			"internal/entity",
			"internal/usecase/interfaces",
			"internal/infrastructure/db",
			"internal/infrastructure/config",
			"internal/delivery/http",
			"pkg",
			"api",
			"web",
			"configs",
		}
	case "hex":
		dirs = []string{
			"cmd/app",
			"internal/core/domain",
			"internal/core/ports/input",
			"internal/core/ports/output",
			"internal/service",
			"internal/adapters/input/http",
			"internal/adapters/output/persistence",
			"web",
			"api",
			"configs",
		}
	case "vertical":
		dirs = []string{
			"cmd/server",
			"internal/features",
			"internal/shared/database",
			"internal/shared/middleware",
			"internal/platform",
			"web",
			"api",
		}
	case "standard":
		dirs = []string{
			"cmd/app",
			"pkg",
			"internal",
			"api",
			"configs",
		}
	case "modular":
		dirs = []string{
			"cmd/app",
			"internal/modules/user",
			"internal/modules/auth",
			"internal/shared/db",
			"internal/shared/logger",
			"web",
			"api",
			"configs",
		}
	default:
		fmt.Printf("Invalid architecture type: %s\n", archType)
		os.Exit(1)
	}

	for _, dir := range dirs {
		path := filepath.Join(basePath, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	// Create a basic main.go with Graceful Shutdown
	mainContent := `package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int    ` + "`" + `yaml:"port"` + "`" + `
		Host string ` + "`" + `yaml:"host"` + "`" + `
	} ` + "`" + `yaml:"server"` + "`" + `
	Database struct {
		Driver string ` + "`" + `yaml:"driver"` + "`" + `
		DSN    string ` + "`" + `yaml:"dsn"` + "`" + `
	} ` + "`" + `yaml:"database"` + "`" + `
}

func main() {
	// Load configuration
	conf := &Config{}
	data, err := os.ReadFile("configs/config.yaml")
	if err == nil {
		yaml.Unmarshal(data, conf)
	}

	fmt.Printf("Project initialized on port %d\n", conf.Server.Port)

	// Graceful Shutdown logic
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	fmt.Println("\nShutting down server gracefully...")
	// Add database connection closures and other cleanup logic here
}
`
	mainPath := filepath.Join(basePath, "cmd", "app", "main.go")
	if archType == "vertical" {
		mainPath = filepath.Join(basePath, "cmd", "server", "main.go")
	}

	if err := os.WriteFile(mainPath, []byte(mainContent), 0644); err != nil {
		fmt.Printf("Error writing main.go: %v\n", err)
		os.Exit(1)
	}

	// Create default config.yaml
	configYaml := `server:
  port: 8080
  host: "0.0.0.0"

database:
  driver: "postgres"
  dsn: "postgres://user:pass@localhost:5432/db?sslmode=disable"
`
	if err := os.WriteFile(filepath.Join(basePath, "configs", "config.yaml"), []byte(configYaml), 0644); err != nil {
		fmt.Printf("Error writing config.yaml: %v\n", err)
		os.Exit(1)
	}

	// Create initial CHANGELOG.md
	changelogContent := `# Changelog

All notable changes to this project will be documented in this file.

## [0.1.0] - ` + strings.Split(fmt.Sprint(os.Args), " ")[0] + `
### Added
- Initial project structure using ` + archType + ` architecture.
- YAML configuration support.
- Graceful shutdown logic.
- Standard linting configuration.
`
	if err := os.WriteFile(filepath.Join(basePath, "CHANGELOG.md"), []byte(changelogContent), 0644); err != nil {
		fmt.Printf("Error writing CHANGELOG.md: %v\n", err)
		os.Exit(1)
	}

	// Create default .golangci.yml
	lintContent := `linters-settings:
  govet:
    check-shadowing: true

linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - typecheck
    - unused
    - gofmt
    - misspell

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0
`
	if err := os.WriteFile(filepath.Join(basePath, ".golangci.yml"), []byte(lintContent), 0644); err != nil {
		fmt.Printf("Error writing .golangci.yml: %v\n", err)
		os.Exit(1)
	}

	// Create .gitignore
	gitignoreContent := `# Binaries
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Dependency directories
vendor/
node_modules/

# IDEs
.vscode/
.idea/

# OS files
.DS_Store
Thumbs.db

# Environment and secrets
.env
*.local.yaml

# Frontend builds
web/dist/
web/.next/
`
	if err := os.WriteFile(filepath.Join(basePath, ".gitignore"), []byte(gitignoreContent), 0644); err != nil {
		fmt.Printf("Error writing .gitignore: %v\n", err)
		os.Exit(1)
	}

	// Create .dockerignore
	dockerignoreContent := `.git
.vscode
.idea
vendor
node_modules
*.log
.env
README.md
CHANGELOG.md
`
	if err := os.WriteFile(filepath.Join(basePath, ".dockerignore"), []byte(dockerignoreContent), 0644); err != nil {
		fmt.Printf("Error writing .dockerignore: %v\n", err)
		os.Exit(1)
	}

	// Create basic README.md
	readmeContent := fmt.Sprintf("# %s\n\nProject generated using the %s pattern.\n\n## How to Run\n1. `make build` to build the binary.\n2. `make run` to run the application.\n3. `make lint` to validate code quality.\n\n## Versioning\nThis project follows SemVer and Conventional Commits.\n", projectName, archType)
	if err := os.WriteFile(filepath.Join(basePath, "README.md"), []byte(readmeContent), 0644); err != nil {
		fmt.Printf("Error writing README.md: %v\n", err)
		os.Exit(1)
	}

	// Create self-documenting Makefile
	makefileContent := `BINARY_NAME=` + projectName + `

.PHONY: help build run test lint clean docker-build

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the application binary
	go build -o bin/$(BINARY_NAME) cmd/app/main.go

run: ## Run the application locally
	go run cmd/app/main.go

test: ## Run unit tests
	go test ./...

lint: ## Run static analysis (linter)
	golangci-lint run

clean: ## Clean binaries and temporary files
	rm -rf bin/
	rm -f .env

docker-build: ## Build the production Docker image
	docker build -t $(BINARY_NAME) .
`
	if archType == "vertical" {
		makefileContent = strings.Replace(makefileContent, "cmd/app/main.go", "cmd/server/main.go", -1)
	}
	if err := os.WriteFile(filepath.Join(basePath, "Makefile"), []byte(makefileContent), 0644); err != nil {
		fmt.Printf("Error writing Makefile: %v\n", err)
		os.Exit(1)
	}

	// Create multi-stage Dockerfile
	appDockerfile := `# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app/main.go

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs
EXPOSE 8080
CMD ["./main"]
`
	if archType == "vertical" {
		appDockerfile = strings.Replace(appDockerfile, "./cmd/app/main.go", "./cmd/server/main.go", -1)
	}
	if err := os.WriteFile(filepath.Join(basePath, "Dockerfile"), []byte(appDockerfile), 0644); err != nil {
		fmt.Printf("Error writing Dockerfile: %v\n", err)
		os.Exit(1)
	}

	// Create basic go.mod
	gomodContent := `module github.com/organization/` + projectName + `

go 1.24

require gopkg.in/yaml.v3 v3.0.1
`
	if err := os.WriteFile(filepath.Join(basePath, "go.mod"), []byte(gomodContent), 0644); err != nil {
		fmt.Printf("Error writing go.mod: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Project generated successfully in /output/" + projectName)
}
