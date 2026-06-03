# Contributing to go-layout

Thank you for your interest in improving the organizational Go standards. This project follows strict technical and professional guidelines.

## Code of Conduct
- Maintain a professional and objective tone in all communications.
- Focus on technical rationale and empirical evidence.

## How to Contribute

### Branching Strategy
- **Direct commits to `main` are strictly prohibited.**
- All changes must be made in a feature or fix branch.
- Changes must be merged via **Pull Request (PR)** after passing CI checks.

### Reporting Issues
- Use the issue tracker to report bugs or suggest enhancements.
- Provide a clear description and steps to reproduce for bugs.

### Suggesting New Patterns
To propose a new architecture or standard:
1. Open an Issue with the prefix `[PROPOSAL]`.
2. Describe the technical benefits and target use case.
3. Provide a directory structure example.

### Development Process
1. Fork the repository.
2. Create a branch: `feat/new-pattern` or `fix/bug-description`.
3. Implement changes ensuring:
    - Markdown files are strictly technical.
    - `generator.go` is updated to support new patterns.
    - All code passes `golangci-lint`.
4. Submit a Pull Request.

## Technical Standards

### Commit Messages
All commits must follow **Conventional Commits**:
- `feat(scope): description`
- `fix(scope): description`
- `docs(scope): description`
- `refactor(scope): description`

### Pull Request Requirements
- The PR must clearly describe the "What" and "Why".
- If the generator was modified, verify it by generating a test project.
- Documentation updates must be objective and free of marketing/praise language.

---
*By contributing, you agree that your contributions will be licensed under the MIT License.*
