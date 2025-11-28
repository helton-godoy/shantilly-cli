# Product Context - Shantilly-CLI TUI

## Project Overview

**Name:** Shantilly-CLI  
**Type:** Terminal User Interface (TUI) Tool  
**Language:** Golang 1.21+  
**Framework:** Charmbracelet (bubbletea, lipgloss, bubbles)  
**Purpose:** Modern alternative to dialog/whiptail for shell scripts  

---

## Mission Statement

Create a modern, cross-platform TUI tool that provides beautiful and functional terminal interfaces for shell scripts, replacing traditional tools like `dialog` and `whiptail` with enhanced user experience and visual appeal.

---

## Core Requirements

### Technical Stack
- **Language:** Golang 1.21+
- **TUI Framework:** Charmbracelet ecosystem
- **Distribution:** Single binary executable
- **Cross-platform:** Linux, macOS, Windows

### Essential Features
1. **Dialog Boxes** - Interactive menus and selections
2. **Forms** - Data collection with validation
3. **Progress Bars** - Visual progress indicators
4. **File Selectors** - File and directory navigation
5. **Confirmations** - Yes/No dialogues

### Target Users
- Shell script developers
- System administrators
- DevOps engineers
- CLI tool creators

---

## Development Framework

This project uses **BMAD-GitHub-Native-Full-Cycle** for autonomous AI-driven development.

### BMAD Personas
1. **Project Manager** - Requirements and planning
2. **Architect** - System design and technical decisions
3. **Developer** - Code implementation in Go
4. **QA** - Testing and quality assurance
5. **Security** - Security analysis and compliance
6. **DevOps** - Build and deployment automation
7. **Release Manager** - Version management and releases

### Development Process
- GitHub-centric workflow
- Micro-commit protocol
- Automated quality gates
- Cross-platform validation

---

## Success Metrics

### Functional Requirements
- [ ] All TUI components working
- [ ] Cross-platform compatibility
- [ ] Binary size < 10MB
- [ ] Startup time < 100ms
- [ ] Memory usage < 50MB

### Quality Requirements
- [ ] 90%+ test coverage
- [ ] Zero security vulnerabilities
- [ ] Documentation complete
- [ ] Performance benchmarks met

### User Experience
- [ ] Intuitive interface
- [ ] Consistent styling
- [ ] Error handling graceful
- [ ] Accessibility features

---

## Architecture Overview

### Expected Structure
```
shantilly-cli/
├── cmd/
│   └── shantilly/
│       └── main.go          # Entry point
├── pkg/
│   ├── tui/
│   │   ├── dialog.go        # Dialog components
│   │   ├── form.go          # Form components
│   │   ├── progress.go      # Progress bars
│   │   └── selector.go      # File selectors
│   └── config/
│       └── config.go        # Configuration
├── go.mod                    # Go modules
├── go.sum                    # Dependencies lock
├── Makefile                  # Build automation
├── README.md                 # Documentation
└── docs/                     # Additional docs
```

### Key Dependencies
```go
// Core Charmbracelet packages
github.com/charmbracelet/bubbletea v0.25.0
github.com/charmbracelet/lipgloss v0.9.1
github.com/charmbracelet/bubbles v0.17.0
```

---

## Target Commands

### Dialog Interface
```bash
shantilly dialog --title "Menu" --options "Option1,Option2,Option3"
```

### Form Interface
```bash
shantilly form --fields "name:text,email:email,age:number"
```

### Progress Interface
```bash
shantilly progress --steps "Build,Test,Deploy" --current 2
```

### File Selection
```bash
shantilly select --path "/home/user" --filter "*.go" --multi
```

### Confirmation
```bash
shantilly confirm --message "Continue?" --default yes
```

---

## Quality Standards

### Code Quality
- Go best practices
- Comprehensive error handling
- Resource management
- Memory efficiency

### Testing
- Unit tests for all components
- Integration tests for workflows
- Cross-platform testing
- Performance benchmarks

### Security
- Input validation
- Safe file operations
- Memory safety
- Dependency scanning

---

## Deployment Strategy

### Build Process
- GitHub Actions CI/CD
- Multi-platform builds
- Automated releases
- Binary distribution

### Release Channels
- GitHub Releases
- Package managers (brew, apt, chocolatey)
- Direct downloads
- Docker images

---

## Future Roadmap

### Version 1.0 (Current)
- Basic TUI components
- Cross-platform support
- Core command set

### Version 1.1
- Advanced styling options
- Theme support
- Plugin system

### Version 2.0
- Web-based TUI
- Remote sessions
- Advanced animations

---

## Project Governance

### Development Methodology
- BMAD autonomous development
- GitHub-centric workflow
- Continuous integration
- Automated quality gates

### Community Guidelines
- Open source contribution
- Issue templates
- PR templates
- Code of conduct

---

**Last Updated:** 2025-11-28T09:50:00Z  
**Version:** 1.0.0  
**Framework:** BMAD-GitHub-Native-Full-Cycle
