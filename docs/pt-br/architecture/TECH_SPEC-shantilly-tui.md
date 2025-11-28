# Technical Specification - Shantilly-CLI TUI

## Architecture Overview

### System Design
Shantilly-CLI é uma ferramenta TUI (Terminal User Interface) moderna construída em Go, utilizando o ecossistema Charmbracelet para criar interfaces terminal elegantes e funcionais.

### Core Components

#### 1. CLI Entry Point (`cmd/shantilly/main.go`)
```go
package main

import (
    "os"
    "github.com/spf13/cobra"
    "github.com/helton-godoy/shantilly-cli/pkg/tui"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "shantilly",
        Short: "Modern TUI for shell scripts",
        Long:  "Shantilly-CLI provides beautiful terminal interfaces",
    }
    
    // Subcommands
    rootCmd.AddCommand(createDialogCmd())
    rootCmd.AddCommand(createFormCmd())
    rootCmd.AddCommand(createProgressCmd())
    rootCmd.AddCommand(createSelectCmd())
    rootCmd.AddCommand(createConfirmCmd())
    
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}
```

#### 2. TUI Package Structure (`pkg/tui/`)

##### Dialog Component (`pkg/tui/dialog.go`)
```go
package tui

import (
    "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type DialogModel struct {
    choices []string
    cursor  int
    selected string
    title   string
}

func (m DialogModel) Init() tea.Cmd { return nil }
func (m DialogModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { /* ... */ }
func (m DialogModel) View() string { /* ... */ }
```

##### Form Component (`pkg/tui/form.go`)
```go
type Field struct {
    Name        string
    Type        string // text, email, number, password
    Required    bool
    Value       string
    Placeholder string
    Validator   func(string) error
}

type FormModel struct {
    fields  []Field
    cursor  int
    focused int
    values  map[string]string
    errors  map[string]string
}
```

##### Progress Component (`pkg/tui/progress.go`)
```go
type ProgressModel struct {
    steps    []string
    current  int
    total    int
    percent  float64
    running  bool
}
```

##### Selector Component (`pkg/tui/selector.go`)
```go
type FileSelectorModel struct {
    path     string
    files    []os.FileInfo
    cursor   int
    selected []string
    filter   string
    multi    bool
}
```

### Dependencies Management

#### Go Modules (`go.mod`)
```go
module github.com/helton-godoy/shantilly-cli

go 1.21

require (
    github.com/charmbracelet/bubbletea v0.25.0
    github.com/charmbracelet/lipgloss v0.9.1
    github.com/charmbracelet/bubbles v0.17.0
    github.com/spf13/cobra v1.8.0
    github.com/spf13/viper v1.17.0
)
```

### CLI Interface Design

#### Command Structure
```bash
shantilly [command] [flags]

Commands:
  dialog     Create interactive dialog boxes
  form       Create data collection forms
  progress   Show progress indicators
  select     File and directory selection
  confirm    Yes/No confirmation dialogs
  help       Help about any command
```

#### Dialog Command
```bash
shantilly dialog [flags]

Flags:
  -t, --title string      Dialog title (required)
  -o, --options strings   Comma-separated options (required)
  -d, --default string    Default selection
  -m, --multi             Allow multiple selections
  -h, --height int        Dialog height (default 10)
  -w, --width int         Dialog width (default 50)
```

#### Form Command
```bash
shantilly form [flags]

Flags:
  -f, --fields strings   Field definitions (required)
  -v, --values strings   Default values
  -t, --title string     Form title
  -w, --width int        Form width (default 60)
```

### Styling System

#### Lipgloss Theme Configuration
```go
var (
    // Base colors
    primaryColor   = lipgloss.Color("#007DCC")
    secondaryColor = lipgloss.Color("#6C757D")
    successColor   = lipgloss.Color("#28A745")
    warningColor   = lipgloss.Color("#FFC107")
    errorColor     = lipgloss.Color("#DC3545")
    
    // Base styles
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(primaryColor).
        MarginBottom(1)
        
    borderStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(primaryColor)
        
    selectedStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("black")).
        Background(primaryColor)
)
```

### Error Handling Strategy

#### Error Types
```go
type ShantillyError struct {
    Code    string
    Message string
    Cause   error
}

const (
    ErrInvalidInput   = "INVALID_INPUT"
    ErrFileNotFound   = "FILE_NOT_FOUND"
    ErrPermission     = "PERMISSION_DENIED"
    ErrNetworkError   = "NETWORK_ERROR"
)
```

#### Graceful Degradation
- Fallback to basic terminal output if TUI fails
- Automatic terminal capability detection
- Safe mode for limited environments

### Performance Considerations

#### Memory Management
- Object pooling for frequently allocated structs
- Efficient string operations
- Minimal external dependencies

#### Rendering Optimization
- Virtual rendering for large lists
- Lazy loading of file lists
- Incremental updates only

### Cross-Platform Compatibility

#### Platform Detection
```go
type Platform int

const (
    Windows Platform = iota
    Linux
    MacOS
)

func DetectPlatform() Platform {
    // Platform-specific detection logic
}
```

#### Terminal Compatibility
- ANSI escape sequence handling
- Windows Console API integration
- Terminal size detection

### Security Considerations

#### Input Validation
- Type validation for all inputs
- Path traversal protection
- Command injection prevention

#### File Operations
- Permission checks before operations
- Safe file path handling
- Temporary file cleanup

### Testing Strategy

#### Unit Tests
- Component-level testing
- Mock external dependencies
- Edge case coverage

#### Integration Tests
- CLI command testing
- Cross-platform validation
- Performance benchmarks

#### E2E Tests
- Full workflow testing
- Terminal interaction simulation
- Error scenario testing

### Build System

#### Makefile Structure
```makefile
.PHONY: build test clean install

# Build for current platform
build:
	go build -o bin/shantilly cmd/shantilly/main.go

# Cross-platform builds
build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/shantilly-linux-amd64 cmd/shantilly/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/shantilly-darwin-amd64 cmd/shantilly/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/shantilly-windows-amd64.exe cmd/shantilly/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install locally
install: build
	cp bin/shantilly /usr/local/bin/
```

### Documentation Structure

#### Code Documentation
- Go doc comments for all public APIs
- Usage examples in code
- Architecture decision records

#### User Documentation
- Command reference
- Tutorial guide
- Troubleshooting guide

### Future Extensibility

#### Plugin System (v1.1)
- Interface-based plugin architecture
- Dynamic plugin loading
- Plugin distribution via go modules

#### Theme System (v1.1)
- YAML-based theme definitions
- Custom color schemes
- Component styling overrides

#### Web Integration (v2.0)
- WebSocket-based remote TUI
- Browser interface support
- Session management

---

**Status:** ✅ **ARCHITECTURE COMPLETE**  
**Next:** Implementation Phase  
**Responsível:** Developer Agent  
**Tech Stack:** Go 1.21+ + Charmbracelet  
**Estimated Implementation:** 3-4 sprints
