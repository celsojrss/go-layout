# Vertical Slice Specification

## Directory Structure
- `internal/features/`: Feature-specific directories containing handler, logic, and data access.
- `internal/shared/`: Shared infrastructure code (DB connections, Middlewares).

## Constraints
- Maintain high cohesion within feature directories.
- Minimize coupling between different feature slices.
- Shared code should be strictly technical and infrastructure-related.
