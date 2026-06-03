# Versioning Standard

## SemVer 2.0.0
- **MAJOR**: Breaking changes.
- **MINOR**: New features (backwards compatible).
- **PATCH**: Bug fixes (backwards compatible).

## Commit Messages
Format: `<type>(<scope>): <description>`

### Types
- `feat`: New feature.
- `fix`: Bug fix.
- `docs`: Documentation.
- `refactor`: Code change without behavior change.
- `chore`: Maintenance tasks.

## SemVer and Commit Mapping

| Commit Type | SemVer Increment | Description |
| :--- | :--- | :--- |
| `feat` | **MINOR** | New backwards-compatible functionality. |
| `fix` | **PATCH** | Backwards-compatible bug fix. |
| `feat!` / `fix!` | **MAJOR** | Breaking changes (API incompatible). |
| `BREAKING CHANGE` | **MAJOR** | Breaking changes (found in commit footer). |
| `docs` / `chore` | **NONE** | No version increment (or internal metadata). |
| `refactor` / `perf`| **PATCH** | Typically a patch increment if it affects stability. |

## Automated Release Tools
- [semantic-release](https://github.com/semantic-release/semantic-release)
- [go-semantic-release](https://github.com/go-semantic-release/go-semantic-release)
