# How to use the Project Scaffolder via Docker

This solution provides a "CLI in a container" that does not require Go to be installed locally.

## 1. Build the Scaffolder
Execute this command inside the `go-layout/scaffolder` folder:
```powershell
docker build -t go-generator -f Dockerfile.generator .
```

## 2. Creating a New Project
To create a new project in your current directory, execute:

### For Clean Architecture:
```powershell
docker run --rm -v ${PWD}:/output go-generator my-clean-project clean
```

### For Hexagonal Architecture:
```powershell
docker run --rm -v ${PWD}:/output go-generator my-hex-project hex
```

### For Vertical Slice:
```powershell
docker run --rm -v ${PWD}:/output go-generator my-vertical-project vertical
```

### For Modular Monolith:
```powershell
docker run --rm -v ${PWD}:/output go-generator my-modular-project modular
```

### For Standard Layout:
```powershell
docker run --rm -v ${PWD}:/output go-generator my-standard-project standard
```

## How it works:
1. `--rm`: Removes the container after execution.
2. `-v ${PWD}:/output`: Maps your current folder to the container's output.
3. `go-generator`: Uses the image built in step 1.
4. `my-project-xxx`: The name of the project directory to be created.
5. `clean|hex|vertical|modular|standard`: The target architecture type.

---
**Technical Note:** This method ensures that everyone in the organization uses the **same Go version** and the **same directory structure**, regardless of their local environment. This represents high-level standardization (Infrastructure as Code for scaffolding).
