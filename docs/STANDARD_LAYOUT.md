# Standard Library Layout Specification (Minimalist)

## Overview
Recommended for small tools, CLI utilities, or shared libraries. It avoids deep nesting and prioritizes simplicity.

## Directory Structure
- `cmd/`: Application entry points.
- `pkg/`: Exportable code (libraries).
- `internal/`: Private utility code.
- `api/`: API definitions.

## Constraints
- Keep logic as close to the root as possible for libraries.
- Use this layout only when business logic complexity is low.
- Prioritize Go's standard library conventions.
