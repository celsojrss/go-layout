# Clean Architecture Specification

## Directory Structure
- `cmd/`: Application entry points.
- `internal/entity/`: Domain models and business rules.
- `internal/usecase/`: Application-specific business rules.
- `internal/delivery/`: Interface adapters (HTTP, gRPC).
- `internal/infrastructure/`: Technical implementations (DB, Cache).

## Constraints
- Dependencies must point towards `internal/entity`.
- No infrastructure details in the `usecase` or `entity` layers.
- Use dependency injection to provide implementations to usecases.
