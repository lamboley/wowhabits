# Wow Habits

## Quick Start

Run `make help` for all available targets. Common workflows:

```
make build               # Build WoWHabits binary
make test                # Run go unit test
make lint                # Run golangci-lint
make format              # Run gofmt and goimports
```

## Codebase Structure

```text
wowhabits/
├── bin/                      # Build output (git-ignored)
├── cmd/
│   └── wowhabits/            # Entry point for the wowhabits binary (Cobra commands)
├── docs/                     # Project documentation
├── pkg/
│   └── api/                  # HTTP adapter: REST routes and handlers (net/http)
└── ui/                       # Web frontend (Vue 3)
```

## Code Standards

- **Changes**: Keep them focused and reviewable; add or update relevant tests
- **Commits**: No `@mentions` or `fixes #...` keywords, no `Co-authored-by:` trailers
- **Packages**: lowercase, single word, match directory.
- **Security**: Be aware of the [Eight Principles of Saltzer and Schroeder](https://web.mit.edu/Saltzer/www/publications/protection/Basic.html), [OWASP TOP10](https://top10.owasp.org/2025/) and [CWE TOP25](https://cwe.mitre.org/top25/archive/2025/2025_cwe_top25.html).

## Documentation

- **[CONTRIBUTING.md](CONTRIBUTING.md)**: Development guidelines
- **[README.md](README.md)**: Project overview
