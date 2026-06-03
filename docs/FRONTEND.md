# Frontend Standard (React)

## Tech Stack
- **Framework**: React 18+ (TypeScript).
- **Build Tool**: Vite.
- **Styling**: Vanilla CSS or Tailwind CSS (Optional).
- **Quality**: ESLint + Prettier.

## Directory Structure
```text
web/
├── src/
│   ├── assets/       # Static assets (images, icons)
│   ├── components/   # Shared UI components (Atomic Design)
│   ├── features/     # Feature-based logic (Hooks, Services, Components)
│   ├── App.tsx       # Main component
│   └── main.tsx      # Entry point
├── public/           # Public assets
├── package.json      # Dependencies and SemVer
└── vite.config.ts    # Vite configuration
```

## Security
- **Scanning**: Gitleaks also scans the `web/` directory.
- **Dependencies**: No sensitive data in environment variables (`.env`) should be committed.
- **XSS/CSRF**: Use standard React security practices (sanitization).

## Integration
The frontend is built into the `dist/` directory and embedded into the Go binary using `go:embed`.
