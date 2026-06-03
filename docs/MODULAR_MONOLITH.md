# Modular Monolith Specification

## Overview
Recommended for large applications that require modularity without the overhead of microservices. Each module is self-contained.

## Directory Structure
- `internal/modules/`: Main container for modules.
    - `<module_name>/`: A complete unit.
        - `domain/`: Entities and logic.
        - `handler/`: Input delivery.
        - `repository/`: Data access.
- `internal/shared/`: Code shared between modules (DB connections, logging).

## Constraints
- Modules should communicate through interfaces or internal events.
- Avoid direct cross-module database access; use service calls or event buses.
- Shared code must be generic and infrastructure-focused.
