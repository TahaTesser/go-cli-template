# Go CLI Template

A modern command-line interface template built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for interactive TUI functionality and [Lipgloss](https://github.com/charmbracelet/lipgloss) for beautiful styling.

## Table of Contents

- [What is this project?](#what-is-this-project)
- [Requirements](#requirements)
- [Quick Start](#quick-start)
- [Usage](#usage)
- [Go Installation Guide](#go-installation-guide)
- [Project Structure](#project-structure)
- [Development](#development)

## What is this project?

This template provides a foundation for creating interactive command-line applications in Go. It demonstrates:

- **Interactive TUI**: Menu navigation with keyboard controls
- **Beautiful styling**: Colors, borders, and typography with Lipgloss
- **CLI framework**: Cobra for command structure and flags
- **Clean architecture**: Organized project structure following Go conventions

Perfect for building tools like file managers, system monitors, configuration utilities, or any CLI that benefits from interactive menus.

## Requirements

- Go 1.21 or later
- Terminal with color support (most modern terminals)

## Quick Start

```bash
# Clone or download this template
git clone <your-repo-url>
cd go-cli-template

# Install dependencies
go mod download

# Build the application
go build -o cli-app

# Run the interactive CLI
./cli-app
```

## Usage

### Interactive Mode
```bash
./cli-app
```

Navigate with:
- `↑/↓` or `j/k` - Move cursor
- `Space/Enter` - Select options
- `q` or `Ctrl+C` - Quit

### Command-line Flags
```bash
./cli-app --help              # Show help
./cli-app --config config.yml # Use config file
./cli-app --verbose           # Enable verbose output
```

## Go Installation Guide

New to Go? Here's how to get started:

### 1. Install Go

**macOS (Homebrew):**
```bash
brew install go
```

**Linux/macOS (Manual):**
```bash
# Download from https://golang.org/dl/
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

**Windows:**
Download installer from [golang.org/dl](https://golang.org/dl/)

### 2. Verify Installation
```bash
go version  # Should show Go version 1.21+
```

### 3. Basic Go Commands
```bash
go mod init my-project     # Initialize new project
go mod download           # Download dependencies
go build                  # Build executable
go run main.go           # Build and run directly
go test ./...            # Run tests
```

## Project Structure

```
.
├── cmd/
│   └── root.go          # CLI commands and flags
├── internal/
│   ├── tui/
│   │   ├── model.go     # Application state
│   │   ├── update.go    # Input handling
│   │   └── view.go      # UI rendering
│   └── styles/
│       └── styles.go    # Colors and styling
├── main.go              # Entry point
├── go.mod               # Dependencies
└── README.md
```

## Development

### Building
```bash
# Development build
go build -o cli-app

# Cross-platform builds
GOOS=linux GOARCH=amd64 go build -o cli-app-linux
GOOS=windows GOARCH=amd64 go build -o cli-app.exe
GOOS=darwin GOARCH=amd64 go build -o cli-app-mac
```

### Adding Features
1. **New commands**: Add to `cmd/root.go`
2. **UI changes**: Modify files in `internal/tui/`
3. **Styling**: Update `internal/styles/styles.go`
4. **Dependencies**: `go get package-name`

### Testing
```bash
go test ./...                # Run all tests
go test -cover ./...         # With coverage
go test -v ./internal/tui/   # Specific package
```

---

Built with ❤️ using [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss)