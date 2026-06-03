# How to use the Project Scaffolder via Docker

This solution provides a "CLI in a container" that does not require Go to be installed locally.

## 1. Build the Scaffolder
Execute this command inside the `go-layout/scaffolder` folder:
```powershell
docker build -t go-generator -f Dockerfile.generator .
```

## 2. Creating a New Project
To create a new project in your current directory, execute:

### Basic Go Project (Backend Only):
```powershell
docker run --rm -v ${PWD}:/output go-generator my-project clean
```

### Full-stack Project (Go + React UI):
Add the `with-ui` flag at the end:
```powershell
docker run --rm -v ${PWD}:/output go-generator my-fullstack-app clean with-ui
```

## 3. Features of Full-stack Projects
When using `with-ui`, the scaffolder:
1.  **Creates a `web/` directory**: Contains a modern React + TypeScript + Vite project.
2.  **Implements `go:embed`**: The Go `main.go` will automatically embed and serve the React `dist` folder.
3.  **Updates the `Makefile`**: Adds a `build-ui` step. Running `make build` will now build the frontend first, then the backend.
4.  **Ready for Production**: The single binary generated will serve the API and the UI simultaneously on port 8080.

## How it works:
1. `--rm`: Removes the container after execution.
2. `-v ${PWD}:/output`: Maps your current folder to the container's output.
3. `go-generator`: Uses the image built in step 1.
4. `my-project-xxx`: The name of the project directory to be created.
5. `clean|hex|vertical|modular|standard`: The target architecture type.
6. `with-ui` (Optional): Includes the React frontend boilerplate.

---
**Technical Note:** This method ensures that everyone in the organization uses the **same Go version** and the **same directory structure**, regardless of their local environment.
