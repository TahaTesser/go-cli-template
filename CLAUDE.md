# Go CLI Template with Bubble Tea & Lipgloss

This template provides a foundation for creating interactive command-line applications in Go using Bubble Tea for TUI (Terminal User Interface) and Lipgloss for styling.

## Project Structure

```
.
├── cmd/
│   └── root.go          # Root command and CLI setup
├── internal/
│   ├── tui/
│   │   ├── model.go     # Bubble Tea model
│   │   ├── update.go    # Update logic
│   │   └── view.go      # View rendering
│   └── styles/
│       └── styles.go    # Lipgloss styles
├── main.go              # Application entry point
├── go.mod
└── go.sum
```

## Key Dependencies

```go
// go.mod
module your-cli-app

go 1.21

require (
    github.com/charmbracelet/bubbletea v0.25.0
    github.com/charmbracelet/lipgloss v0.9.1
    github.com/spf13/cobra v1.8.0
)
```

## Basic CLI Setup Commands

### Initialize project
```bash
go mod init your-cli-app
go get github.com/charmbracelet/bubbletea@latest
go get github.com/charmbracelet/lipgloss@latest
go get github.com/spf13/cobra@latest
```

### Build and run
```bash
go build -o cli-app
./cli-app --help
```

### Development
```bash
go run main.go
```

## Core Components

### 1. Main Entry Point (main.go)
- Initialize Cobra root command
- Set up global flags and configuration
- Handle application lifecycle

### 2. Bubble Tea Model (internal/tui/model.go)
- Define application state
- Implement bubbletea.Model interface
- Handle initialization and cleanup

### 3. Update Logic (internal/tui/update.go)
- Process keyboard input and messages
- Update application state
- Handle navigation and user interactions

### 4. View Rendering (internal/tui/view.go)
- Render UI components using Lipgloss
- Apply styling and layout
- Display application state

### 5. Styles (internal/styles/styles.go)
- Define consistent styling with Lipgloss
- Color schemes and typography
- Reusable UI components

## Common CLI Patterns

### Basic Flags
```go
// String flag
rootCmd.PersistentFlags().StringP("config", "c", "", "config file path")

// Boolean flag
rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

// Integer flag
rootCmd.PersistentFlags().IntP("port", "p", 8080, "server port")
```

### Subcommands
```go
// Add subcommand to root
rootCmd.AddCommand(serveCmd)
rootCmd.AddCommand(initCmd)
rootCmd.AddCommand(versionCmd)
```

### Configuration
```go
// Bind flags to viper
viper.BindPFlags(rootCmd.PersistentFlags())

// Environment variable support
viper.SetEnvPrefix("CLI_APP")
viper.AutomaticEnv()
```

## Bubble Tea Patterns

### Model Structure
```go
type Model struct {
    state     AppState
    cursor    int
    choices   []string
    selected  map[int]struct{}
    err       error
}
```

### Key Bindings
```go
switch msg := msg.(type) {
case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "q":
        return m, tea.Quit
    case "up", "k":
        // Move cursor up
    case "down", "j":
        // Move cursor down
    case "enter", " ":
        // Select item
    }
}
```

## Lipgloss Styling

### Color Schemes
```go
var (
    primaryColor   = lipgloss.Color("#7D56F4")
    secondaryColor = lipgloss.Color("#F25D94")
    accentColor    = lipgloss.Color("#04B575")
)
```

### Component Styles
```go
var (
    titleStyle = lipgloss.NewStyle().
        Foreground(primaryColor).
        Bold(true).
        Padding(1, 2)
    
    itemStyle = lipgloss.NewStyle().
        PaddingLeft(2).
        Foreground(lipgloss.Color("#FAFAFA"))
)
```

## Testing Commands

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestModelUpdate ./internal/tui/

# Benchmark tests
go test -bench=. ./...
```

## Build Commands

```bash
# Build for current platform
go build -o bin/cli-app

# Cross-compile for multiple platforms
GOOS=linux GOARCH=amd64 go build -o bin/cli-app-linux-amd64
GOOS=windows GOARCH=amd64 go build -o bin/cli-app-windows-amd64.exe
GOOS=darwin GOARCH=amd64 go build -o bin/cli-app-darwin-amd64

# Build with version info
go build -ldflags="-X 'main.Version=v1.0.0'" -o bin/cli-app
```

## Linting and Formatting

```bash
# Format code
go fmt ./...

# Run linter (requires golangci-lint)
golangci-lint run

# Run vet
go vet ./...

# Generate documentation
go doc -all
```

## Common TUI Features

- Interactive menus and lists
- Progress bars and spinners
- Form inputs and validation
- Multi-pane layouts
- Keyboard navigation
- Mouse support (optional)
- Color and styling
- Responsive design

## Example Usage Patterns

```bash
# Interactive mode
./cli-app

# Direct command execution
./cli-app serve --port 8080

# Configuration file
./cli-app --config ./config.yaml

# Verbose output
./cli-app --verbose command

# Help and documentation
./cli-app --help
./cli-app command --help
```