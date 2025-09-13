# Repository Guidelines

## Project Structure & Modules
- `main.go`: entrypoint; wires Cobra `cmd` package.
- `cmd/`: CLI surface (commands, flags). Start in `cmd/root.go`.
- `internal/tui/`: Bubble Tea model: `model.go` (state), `update.go` (inputs), `view.go` (rendering).
- `internal/styles/`: Lipgloss styles shared across views.
- Build artifact: `cli-app` (created by `go build`).

## Build, Test, and Dev Commands
- Install deps: `go mod download`.
- Run locally: `go run .` or `./cli-app` after build.
- Build: `go build -o cli-app` (add `GOOS/GOARCH` for cross‑builds).
- Lint/format: `gofmt -s -w .` and `go vet ./...`.
- Tests: `go test ./...` and `go test -cover ./...`.

## Coding Style & Naming
- Use `gofmt` defaults (tabs, standard spacing). No manual alignment.
- Package names: short, lowercase (`tui`, `styles`). Files: `snake_case.go`.
- Exported identifiers use `CamelCase`; unexported use `camelCase`.
- Add GoDoc comments for exported types/functions (complete sentences).
- Keep `cmd/` thin; place logic in `internal/` packages.

## Testing Guidelines
- Use the standard `testing` package.
- File naming: `*_test.go`; test funcs: `TestXxx(t *testing.T)`.
- Prefer small, deterministic tests at the package level (e.g., `internal/tui`).
- Run: `go test -v ./...`; aim to cover state transitions in `update.go`.

## Commit & Pull Requests
- Commit messages: follow Conventional Commits (e.g., `feat: add help flag`, `fix: handle window resize`).
- PRs must include: brief summary, rationale, and testing notes; link issues with `Fixes #<id>` when applicable.
- Screenshots/GIFs of the TUI are encouraged for UI changes.
- Keep PRs focused; update README or in‑code docs when behavior changes.

## Tips & Conventions
- Flags live in `cmd/root.go`; pass config into models rather than globals.
- Styling changes belong in `internal/styles/styles.go` to keep views clean.
- For interactive flows, keep `Model` immutable in intent: return updated copies from `Update`.
