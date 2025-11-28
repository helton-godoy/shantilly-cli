package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Theme represents a visual theme for TUI components
type Theme struct {
	Name       string
	Primary    lipgloss.Color
	Secondary  lipgloss.Color
	Success    lipgloss.Color
	Warning    lipgloss.Color
	Error      lipgloss.Color
	Background lipgloss.Color
	Foreground lipgloss.Color
	Border     lipgloss.Color
}

// Predefined themes
var (
	// DefaultTheme - Professional blue theme
	DefaultTheme = Theme{
		Name:       "default",
		Primary:    "#007DCC",
		Secondary:  "#6C757D",
		Success:    "#28A745",
		Warning:    "#FFC107",
		Error:      "#DC3545",
		Background: "#FFFFFF",
		Foreground: "#212529",
		Border:     "#007DCC",
	}

	// DarkTheme - Dark mode theme
	DarkTheme = Theme{
		Name:       "dark",
		Primary:    "#61DAFB",
		Secondary:  "#8892B0",
		Success:    "#64FFDA",
		Warning:    "#FFD700",
		Error:      "#FF6B6B",
		Background: "#0A0E27",
		Foreground: "#CCD6F6",
		Border:     "#61DAFB",
	}

	// GitHubTheme - GitHub inspired theme
	GitHubTheme = Theme{
		Name:       "github",
		Primary:    "#0969DA",
		Secondary:  "#656D76",
		Success:    "#1F883D",
		Warning:    "#9A6700",
		Error:      "#CF222E",
		Background: "#FFFFFF",
		Foreground: "#24292F",
		Border:     "#D0D7DE",
	}

	// CyberpunkTheme - Neon cyberpunk theme
	CyberpunkTheme = Theme{
		Name:       "cyberpunk",
		Primary:    "#00FFFF",
		Secondary:  "#FF00FF",
		Success:    "#00FF00",
		Warning:    "#FFFF00",
		Error:      "#FF0080",
		Background: "#0A0A0A",
		Foreground: "#00FFFF",
		Border:     "#FF00FF",
	}
)

// GetTheme returns a theme by name
func GetTheme(name string) Theme {
	switch name {
	case "dark":
		return DarkTheme
	case "github":
		return GitHubTheme
	case "cyberpunk":
		return CyberpunkTheme
	default:
		return DefaultTheme
	}
}

// GetThemeNames returns all available theme names
func GetThemeNames() []string {
	return []string{"default", "dark", "github", "cyberpunk"}
}

// StyleBuilder creates styled components based on theme
type StyleBuilder struct {
	theme Theme
}

// NewStyleBuilder creates a new style builder with the given theme
func NewStyleBuilder(theme Theme) *StyleBuilder {
	return &StyleBuilder{theme: theme}
}

// Title creates a title style
func (sb *StyleBuilder) Title() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(sb.theme.Primary).
		MarginBottom(1)
}

// Border creates a border style
func (sb *StyleBuilder) Border() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(sb.theme.Border).
		Padding(0, 1)
}

// Selected creates a selected item style
func (sb *StyleBuilder) Selected() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(sb.theme.Background).
		Background(sb.theme.Primary)
}

// Normal creates a normal text style
func (sb *StyleBuilder) Normal() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(sb.theme.Secondary)
}

// Success creates a success style
func (sb *StyleBuilder) Success() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(sb.theme.Success)
}

// Warning creates a warning style
func (sb *StyleBuilder) Warning() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(sb.theme.Warning)
}

// Error creates an error style
func (sb *StyleBuilder) Error() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(sb.theme.Error)
}

// Info creates an info style
func (sb *StyleBuilder) Info() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(sb.theme.Secondary)
}

// Progress creates a progress bar style
func (sb *StyleBuilder) Progress() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(sb.theme.Background).
		Foreground(sb.theme.Primary)
}

// ProgressComplete creates a completed progress bar style
func (sb *StyleBuilder) ProgressComplete() lipgloss.Style {
	return lipgloss.NewStyle().
		Background(sb.theme.Primary).
		Foreground(sb.theme.Background)
}
