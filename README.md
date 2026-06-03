# Go Foundation - Organizational Standards

This repository (or directory) contains the technical foundation for all Go projects within the organization. It is divided into two main parts:

## 1. Standards Documentation (`/docs`)
Architectural and versioning definitions:
- **[Clean Architecture](./docs/CLEAN_ARCHITECTURE.md)**: For complex, long-term systems.
- **[Hexagonal Architecture](./docs/HEXAGONAL_ARCHITECTURE.md)**: For systems with multiple external integrations.
- **[Vertical Slice](./docs/VERTICAL_SLICE.md)**: For agile, feature-based development.
- **[Modular Monolith](./docs/MODULAR_MONOLITH.md)**: For large systems requiring module independence.
- **[Standard Layout](./docs/STANDARD_LAYOUT.md)**: For minimalist libraries and CLI tools.
- **[Versioning & Commits](./docs/VERSIONING.md)**: SemVer and Conventional Commits standards.
- **[Observability](./docs/OBSERVABILITY.md)**: Health checks, logging, and graceful shutdown standards.
- **[Agent Instructions](./docs/AGENTS.md)**: Guidelines for AI agents working on the project.

## 2. Scaffolding Tool (`/scaffolder`)
Automated Docker-based tool to generate new projects following the standards above.
- **[How to use the Scaffolder](./scaffolder/DOCKER_USAGE.md)**

---
*Keep these standards updated and versioned to ensure continuous technical quality across the team.*
