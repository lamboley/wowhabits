# AI Agent Guide for WoWHabits Development

## Persona

This guide is for AI coding assistants (for example Claude, ChatGPT, Copilot). Use it for context on codebase layout, architecture, and code standards when helping **contributors to [lamboley/wowhabits](https://github.com/lamboley/wowhabits)**.

## Project Overview

**WoW Habits** is a gamified habit-tracking app that uses World of Warcraft characters to represent real-life progress.

## Quick Start

```bash
# Lint and Build
make lint-all           # Lint and check formatting (reports only)
golangci-lint fmt       # Fix formatting
make build-all          # Build all binaries

# Development tools
make install-all        # Install linters and dev tools
```

## Codebase Structure

```text
wowhabits/
├── bin/                      # Build output (git-ignored)
├── cmd/wowhabits/            # Entry point for the wowhabits binary
├── docs/                     # Project documentation
├── frontend/                 # Web frontend (Vue 3)
└── internal/
    └── api/                  # HTTP REST API server
```

## Architecture

The backend follows a hexagonal (ports and adapters) architecture, see [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md).

- **Domain**: entities and business rules. Plain Go: no Gin, HTTP, SQL or framework imports.
- **Ports**: interfaces the domain needs from the outside (e.g. a repository), defined next to the code that uses them.
- **Adapters**: implementations of the ports (database, email) and entry points that call into the domain. `internal/api` is the HTTP adapter.

Rules:

- Dependencies point inward only: adapters → domain, never the reverse.
- Keep Gin handlers thin: parse the request, call the domain, write the response. Business logic in a handler belongs in the domain.
- `cmd/wowhabits` wires adapters together; it holds no business logic.

## Code Standards

**Official Documentation**: [CONTRIBUTING.md](CONTRIBUTING.md)

- **Formatter**: `gofumpt` (via `golangci-lint`, configured in `.golangci.yml`)
- **Validation**: All PRs must pass `make lint-all`
- **Changes**: Keep them focused and reviewable; add or update relevant tests
- **Commits**: No `@mentions` or `fixes #...` keywords, no `Co-authored-by:` trailers

## Key Libraries

- **[Cobra](https://github.com/spf13/cobra)**: CLI framework used for `cmd/wowhabits` commands
- **[Gin](https://github.com/gin-gonic/gin)**: HTTP framework for the REST API in `internal/api` (routing, JSON, request binding and validation)
- **[Go](https://go.dev)**: Programming language

## Documentation

- **[CONTRIBUTING.md](CONTRIBUTING.md)**: Development guidelines
- **[docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)**: Hexagonal architecture
- **[docs/CODE_STRUCTURE.md](docs/CODE_STRUCTURE.md)**: Detailed codebase structure
- **[README.md](README.md)**: Project overview
- **[SECURITY.md](SECURITY.md)**: How to report vulnerabilities (never in a public issue)
