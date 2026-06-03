# Hexagonal Architecture Specification

## Directory Structure
- `internal/core/domain/`: Domain logic and entities.
- `internal/core/ports/`: Inbound and Outbound interfaces.
- `internal/adapters/input/`: Driving adapters (HTTP, CLI).
- `internal/adapters/output/`: Driven adapters (SQL, SMTP).

## Constraints
- The `core` must not depend on any adapter.
- Communication between adapters and core must occur through `ports`.
- Use DTOs for data transfer between layers.
