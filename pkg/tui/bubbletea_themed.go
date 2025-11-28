package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// ThemedDialogModel represents a themed dialog box
type ThemedDialogModel struct {
	choices   []string
	cursor    int
	selected  []string
	title     string
	multi     bool
	confirmed bool
	quitting  bool
	theme     Theme
	styles    *StyleBuilder
}

// ThemedDialogOptions contains configuration for themed dialog
type ThemedDialogOptions struct {
	Title   string
	Choices []string
	Default string
	Multi   bool
	Theme   string
}

// NewThemedDialogModel creates a new themed dialog model
func NewThemedDialogModel(options ThemedDialogOptions) ThemedDialogModel {
	theme := GetTheme(options.Theme)
	styles := NewStyleBuilder(theme)

	// Set cursor to default value if provided
	cursor := 0
	if options.Default != "" {
		for i, choice := range options.Choices {
			if choice == options.Default {
				cursor = i
				break
			}
		}
	}

	return ThemedDialogModel{
		choices:  options.Choices,
		cursor:   cursor,
		title:    options.Title,
		multi:    options.Multi,
		selected: []string{},
		theme:    theme,
		styles:   styles,
	}
}

// Init initializes the themed dialog model
func (m ThemedDialogModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m ThemedDialogModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

// View renders the themed dialog
func (m ThemedDialogModel) View() string {
	if m.quitting {
		return ""
	}

	var content strings.Builder

	// Title
	content.WriteString(m.styles.Title().Render(m.title))
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
			line = m.styles.Selected().Render(line)
		} else {
			line = m.styles.Normal().Render(line)
		}

		content.WriteString(line)
		content.WriteString("\n")
	}

	// Instructions
	content.WriteString("\n")
	if m.multi {
		content.WriteString(m.styles.Info().Render("↑↓ Navigate • Space Toggle • Enter Confirm • Esc Quit"))
	} else {
		content.WriteString(m.styles.Info().Render("↑↓ Navigate • Enter Select • Esc Quit"))
	}

	return m.styles.Border().Render(content.String())
}

// ThemedShowDialog displays a themed dialog and returns the selection
func ThemedShowDialog(title string, choices []string, defaultVal string, multi bool, themeName string) ([]string, error) {
	if len(choices) == 0 {
		return nil, fmt.Errorf("no choices provided")
	}

	options := ThemedDialogOptions{
		Title:   title,
		Choices: choices,
		Default: defaultVal,
		Multi:   multi,
		Theme:   themeName,
	}

	model := NewThemedDialogModel(options)
	p := tea.NewProgram(model)

	finalModel, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run themed dialog: %w", err)
	}

	dialogModel := finalModel.(ThemedDialogModel)
	if dialogModel.quitting {
		return nil, fmt.Errorf("user cancelled")
	}

	return dialogModel.selected, nil
}
