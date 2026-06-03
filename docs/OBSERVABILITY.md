# Quality and Observability Standards

## Code Quality (Linting)
All projects must pass the `golangci-lint` check. The standard configuration is provided in the `.golangci.yml` file at the project root.
- **Rule:** Never merge code with linting errors.
- **Tools:** Use `make lint` to run the analysis locally.

## Observability
### Health Checks
Every web-based project must expose:
- `GET /healthz`: Liveness probe (is the app running?).
- `GET /readyz`: Readiness probe (is the app ready to receive traffic?).

### Graceful Shutdown
Applications must listen for `os.Interrupt` and `syscall.SIGTERM`. When received, the app must:
1. Stop accepting new requests.
2. Finish processing active requests (with timeout).
3. Close database connections and file handles.
4. Exit with code 0.

### Structured Logging
Use structured logging (JSON) in production for better searchability in tools like ELK or Grafana Loki.
