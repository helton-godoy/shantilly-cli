package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BubbleTeaDialogModel represents a dialog box state
type BubbleTeaDialogModel struct {
	choices   []string
	cursor    int
	selected  []string
	title     string
	multi     bool
	confirmed bool
	quitting  bool
}

// BubbleTeaDialogOptions contains configuration for dialog
type BubbleTeaDialogOptions struct {
	Title   string
	Choices []string
	Default string
	Multi   bool
}

// Styles for dialog
var (
	dialogTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#007DCC")).
				MarginBottom(1)

	dialogBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#007DCC")).
				Padding(0, 1)

	dialogSelectedStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("black")).
				Background(lipgloss.Color("#007DCC"))

	dialogNormalStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#6C757D"))
)

// Init initializes the dialog model
func (m BubbleTeaDialogModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m BubbleTeaDialogModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit

		case tea.KeyEnter:
			if m.multi {
				// Toggle selection for multi-select
				for i, choice := range m.choices {
					if i == m.cursor {
						found := false
						for _, sel := range m.selected {
							if sel == choice {
								found = true
								break
							}
						}
						if found {
							m.selected = removeString(m.selected, choice)
						} else {
							m.selected = append(m.selected, choice)
						}
					}
				}
			} else {
				// Single selection - confirm and quit
				m.selected = []string{m.choices[m.cursor]}
				m.confirmed = true
				return m, tea.Quit
			}

		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}

		case tea.KeyDown:
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case tea.KeySpace:
			if m.multi {
				// Toggle selection with spacebar
				choice := m.choices[m.cursor]
				found := false
				for _, sel := range m.selected {
					if sel == choice {
						found = true
						break
					}
				}
				if found {
					m.selected = removeString(m.selected, choice)
				} else {
					m.selected = append(m.selected, choice)
				}
			}
		}
	}

	return m, nil
}

// View renders the dialog
func (m BubbleTeaDialogModel) View() string {
	if m.quitting {
		return ""
	}

	var content strings.Builder

	// Title
	content.WriteString(dialogTitleStyle.Render(m.title))
	content.WriteString("\n")

	// Choices
	for i, choice := range m.choices {
		cursor := " "
		if i == m.cursor {
			cursor = "❯"
		}

		// Check if selected (for multi-select)
		selected := ""
		if m.multi {
			for _, sel := range m.selected {
				if sel == choice {
					selected = "✓ "
					break
				}
			}
		}

		// Apply styling
		line := fmt.Sprintf("%s %s%s", cursor, selected, choice)
		if i == m.cursor {
			line = dialogSelectedStyle.Render(line)
		} else {
			line = dialogNormalStyle.Render(line)
		}

		content.WriteString(line)
		content.WriteString("\n")
	}

	// Instructions
	content.WriteString("\n")
	if m.multi {
		content.WriteString(dialogNormalStyle.Render("↑↓ Navigate • Space Toggle • Enter Confirm • Esc Quit"))
	} else {
		content.WriteString(dialogNormalStyle.Render("↑↓ Navigate • Enter Select • Esc Quit"))
	}

	return dialogBorderStyle.Render(content.String())
}

// BubbleTeaShowDialog displays a BubbleTea dialog and returns the selection
func BubbleTeaShowDialog(title string, choices []string, defaultVal string, multi bool) ([]string, error) {
	if len(choices) == 0 {
		return nil, fmt.Errorf("no choices provided")
	}

	// Set cursor to default value if provided
	cursor := 0
	if defaultVal != "" {
		for i, choice := range choices {
			if choice == defaultVal {
				cursor = i
				break
			}
		}
	}

	model := BubbleTeaDialogModel{
		choices:  choices,
		cursor:   cursor,
		title:    title,
		multi:    multi,
		selected: []string{},
	}

	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run dialog: %w", err)
	}

	dialogModel := finalModel.(BubbleTeaDialogModel)
	if dialogModel.quitting {
		return nil, fmt.Errorf("user cancelled")
	}

	return dialogModel.selected, nil
}
