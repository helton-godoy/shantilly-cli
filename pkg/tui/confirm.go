package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ConfirmModel represents a confirmation dialog state
type ConfirmModel struct {
	message    string
	defaultVal string
	confirmed  bool
	quitting   bool
}

// ConfirmOptions contains configuration for confirmation
type ConfirmOptions struct {
	Message string
	Default string
}

var (
	confirmTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#007DCC"))

	messageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#212529"))

	yesStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("black")).
			Background(lipgloss.Color("#28A745"))

	noStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("black")).
		Background(lipgloss.Color("#DC3545"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C757D"))

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#007DCC"))
)

// Init initializes the confirm model
func (m ConfirmModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m ConfirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit

		case tea.KeyEnter:
			m.confirmed = true
			return m, tea.Quit

		case tea.KeyY, tea.Keyy:
			m.confirmed = true
			return m, tea.Quit

		case tea.KeyN, tea.Keyn:
			m.confirmed = false
			return m, tea.Quit

		case tea.KeyLeft, tea.KeyUp, tea.KeyH, tea.KeyK:
			m.confirmed = false
			return m, tea.Quit

		case tea.KeyRight, tea.KeyDown, tea.KeyL, tea.KeyJ:
			m.confirmed = true
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the confirmation dialog
func (m ConfirmModel) View() string {
	if m.quitting {
		return ""
	}

	var content strings.Builder

	// Title
	content.WriteString(confirmTitleStyle.Render("Confirm"))
	content.WriteString("\n\n")

	// Message
	content.WriteString(messageStyle.Render(m.message))
	content.WriteString("\n\n")

	// Options
	content.WriteString(strings.Repeat(" ", 10))
	content.WriteString(yesStyle.Render("  Yes  "))
	content.WriteString(strings.Repeat(" ", 4))
	content.WriteString(noStyle.Render("  No  "))
	content.WriteString("\n\n")

	// Default indicator
	if m.defaultVal != "" {
		defaultText := fmt.Sprintf("Default: %s", strings.Title(m.defaultVal))
		content.WriteString(infoStyle.Render(defaultText))
		content.WriteString("\n\n")
	}

	// Instructions
	content.WriteString(infoStyle.Render("Y/←/↑ = Yes • N/→/↓ = No • Enter = Confirm • Esc = Cancel"))

	return borderStyle.Render(content.String())
}

// ShowConfirm displays a confirmation dialog
func ShowConfirm(message, defaultVal string) (string, error) {
	if message == "" {
		return "", fmt.Errorf("message cannot be empty")
	}

	// Set default confirmation
	confirmed := true
	if defaultVal == "no" || defaultVal == "n" {
		confirmed = false
	}

	model := ConfirmModel{
		message:    message,
		defaultVal: defaultVal,
		confirmed:  confirmed,
	}

	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run confirmation: %w", err)
	}

	confirmModel := finalModel.(ConfirmModel)
	if confirmModel.quitting {
		return "", fmt.Errorf("user cancelled")
	}

	if confirmModel.confirmed {
		return "yes", nil
	}
	return "no", nil
}
