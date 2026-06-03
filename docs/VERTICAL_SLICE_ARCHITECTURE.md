# Go Vertical Slice Architecture Standard

## Overview
Vertical Slice Architecture focuses on features rather than layers. Each "slice" represents a specific business capability and contains all the code necessary for that feature, from the API handler to the data access.

## Directory Structure
```text
.
├── cmd/                         # Entry points
├── internal/
│   ├── features/                # Domain-driven feature slices
│   │   ├── create_user/         # Everything related to creating a user
│   │   │   ├── handler.go       # HTTP/Entry point
│   │   │   ├── request.go       # DTOs
│   │   │   ├── service.go       # Business Logic
│   │   │   └── repository.go    # Data Access
│   │   └── get_order/           # Everything related to getting an order
│   ├── shared/                  # Common code used across slices
│   │   ├── database/            # DB Connection setup
│   │   ├── middleware/          # Auth, Logging
│   │   └── response/            # Standard JSON wrappers
│   └── platform/                # Low-level utilities
├── web/                         # Frontend (Optionally organized by features)
├── api/                         # API Definitions
└── deployments/                 # Infrastructure as Code
```

## Key Principles

### 1. High Cohesion
All code that changes together for a feature should live together. No more jumping between 5 folders to add a field to a "User".

### 2. Reduced Abstraction
Only abstract what is truly shared. If a feature needs a specific SQL query, write it directly in that feature's repository file.

### 3. Feature Independence
A change in the `create_user` slice should have zero impact on the `get_order` slice.

## Frontend Integration
- Frontend in `/web`.
- Ideally, the frontend structure should mirror the backend features (Feature-based folder structure in React/Vue).
- This allows a "Full-stack Slice" approach where a developer works on both Go and JS in the same logical area.

## Best Practices
- **Shared Folder:** Keep the `shared` folder as lean as possible. If code is only used by two slices, consider duplicating or keeping it in one until a third slice needs it.
- **Refactoring:** When a slice becomes too large, split it into smaller sub-features.
- **Symmetry:** Aim for a consistent structure within each feature folder to maintain predictability.
