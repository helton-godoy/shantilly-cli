# Shantilly-CLI

Modern TUI (Terminal User Interface) tool for shell scripts - Alternative to dialog/whiptail built with Golang + Charmbracelet.

## 🎯 Purpose

Create a modern, cross-platform terminal interface for shell scripts, replacing traditional tools like `dialog` and `whiptail` with a more user-friendly and visually appealing experience.

## 🛠️ Tech Stack

- **Language:** Golang 1.21+
- **TUI Framework:** Charmbracelet (bubbletea, lipgloss, bubbles)
- **Target:** Cross-platform (Linux, macOS, Windows)
- **Distribution:** Single binary executable

## 📋 Features

- **Dialog Boxes** - Interactive menus and selections
- **Forms** - Data collection with validation
- **Progress Bars** - Visual progress indicators  
- **File Selectors** - File and directory navigation
- **Confirmations** - Yes/No dialogues
- **Cross-platform** - Works everywhere

## 🚀 Quick Start

```bash
# Install (after build)
go install ./cmd/shantilly

# Basic usage
shantilly dialog --title "Menu" --options "Option1,Option2,Option3"
shantilly form --fields "name:text,email:email,age:number"
shantilly progress --steps "Build,Test,Deploy" --current 2
```

## 📁 Project Structure

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
└── README.md                 # This file
```

## 🏗️ Development

This project is developed using the **BMAD-GitHub-Native-Full-Cycle** framework for autonomous AI-driven development.

### BMAD Workflow

The development follows the BMAD (Breakthrough Method for Agile AI-Driven Development) methodology with 7 specialized personas:

1. **Project Manager** - Requirements and planning
2. **Architect** - System design and technical decisions  
3. **Developer** - Code implementation
4. **QA** - Testing and quality assurance
5. **Security** - Security analysis and compliance
6. **DevOps** - Infrastructure and deployment
7. **Release Manager** - Version management and releases

## 📄 License

MIT License - see LICENSE file for details.

## 🤝 Contributing

Contributions welcome! Please read the contributing guidelines and submit pull requests.

---

**Built with ❤️ using [BMAD Framework](https://github.com/bmad-code-org/BMAD-METHOD)**