package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Field represents a form field
type Field struct {
	Name        string
	Type        string // text, email, number, password
	Required    bool
	Value       string
	Placeholder string
	Validator   func(string) error
}

// FormModel represents a form state
type FormModel struct {
	fields   []Field
	cursor   int
	focused  int
	values   map[string]string
	errors   map[string]string
	quitting bool
}

// FormOptions contains configuration for form
type FormOptions struct {
	Title  string
	Fields []string
	Values []string
}

var (
	fieldTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#007DCC"))

	inputStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#F8F9FA")).
			Foreground(lipgloss.Color("#212529")).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#007DCC"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#DC3545")).
			Italic(true)

	focusedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#E3F2FD")).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#2196F3"))
)

// Init initializes the form model
func (m FormModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m FormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit

		case tea.KeyEnter:
			// Validate all fields
			if m.validateForm() {
				return m, tea.Quit
			}

		case tea.KeyUp, tea.KeyShiftTab, tea.KeyK:
			if m.cursor > 0 {
				m.cursor--
			}

		case tea.KeyDown, tea.KeyTab, tea.KeyJ:
			if m.cursor < len(m.fields)-1 {
				m.cursor++
			}

		case tea.KeyBackspace:
			if len(m.values[m.fields[m.cursor].Name]) > 0 {
				current := m.values[m.fields[m.cursor].Name]
				m.values[m.fields[m.cursor].Name] = current[:len(current)-1]
				m.validateField(m.fields[m.cursor])
			}

		default:
			// Handle character input
			if len(msg.String()) == 1 {
				field := m.fields[m.cursor]
				current := m.values[field.Name]
				m.values[field.Name] = current + msg.String()
				m.validateField(field)
			}
		}
	}

	return m, nil
}

// View renders the form
func (m FormModel) View() string {
	if m.quitting {
		return ""
	}

	var content strings.Builder

	// Title
	if m.fields[0].Type != "" { // Check if we have actual fields
		content.WriteString(fieldTitleStyle.Render("Form"))
		content.WriteString("\n\n")
	}

	// Fields
	for i, field := range m.fields {
		// Field label
		label := field.Name
		if field.Required {
			label += " *"
		}
		if i == m.cursor {
			label = fieldTitleStyle.Render(label)
		}
		content.WriteString(label)

		// Input field
		value := m.values[field.Name]
		if field.Type == "password" {
			value = strings.Repeat("•", len(value))
		} else if value == "" && field.Placeholder != "" {
			value = field.Placeholder
		}

		inputWidth := 40
		if len(value) > inputWidth {
			value = value[:inputWidth]
		} else {
			value += strings.Repeat(" ", inputWidth-len(value))
		}

		if i == m.cursor {
			content.WriteString(focusedStyle.Render(" " + value + " "))
		} else {
			content.WriteString(inputStyle.Render(" " + value + " "))
		}

		// Error message
		if err, exists := m.errors[field.Name]; exists && err != "" {
			content.WriteString("\n")
			content.WriteString(errorStyle.Render("  " + err))
		}

		content.WriteString("\n\n")
	}

	// Instructions
	content.WriteString(fieldTitleStyle.Render("Instructions"))
	content.WriteString("\n")
	content.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#6C757D")).Render("↑↓ Navigate • Type to input • Enter Submit • Esc Cancel"))

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#007DCC")).
		Padding(1, 2).
		Render(content.String())
}

// validateField validates a single field
func (m *FormModel) validateField(field Field) {
	value := m.values[field.Name]

	// Required validation
	if field.Required && value == "" {
		m.errors[field.Name] = "This field is required"
		return
	}

	// Type validation
	switch field.Type {
	case "email":
		if !strings.Contains(value, "@") || !strings.Contains(value, ".") {
			m.errors[field.Name] = "Invalid email format"
			return
		}
	case "number":
		if value != "" && !isNumeric(value) {
			m.errors[field.Name] = "Must be a number"
			return
		}
	}

	// Custom validator
	if field.Validator != nil {
		if err := field.Validator(value); err != nil {
			m.errors[field.Name] = err.Error()
			return
		}
	}

	// Clear error if validation passes
	delete(m.errors, field.Name)
}

// validateForm validates all fields
func (m *FormModel) validateForm() bool {
	for _, field := range m.fields {
		m.validateField(field)
		if _, exists := m.errors[field.Name]; exists && m.errors[field.Name] != "" {
			return false
		}
	}
	return true
}

// ShowForm displays a form and returns the field values
func ShowForm(title string, fieldDefs []string, defaultValues []string) (map[string]string, error) {
	if len(fieldDefs) == 0 {
		return nil, fmt.Errorf("no fields provided")
	}

	fields := make([]Field, len(fieldDefs))
	values := make(map[string]string)

	for i, fieldDef := range fieldDefs {
		parts := strings.Split(fieldDef, ":")
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid field definition: %s", fieldDef)
		}

		field := Field{
			Name: parts[0],
			Type: parts[1],
		}

		if len(parts) > 2 && parts[2] == "required" {
			field.Required = true
		}

		fields[i] = field

		// Set default value if provided
		if i < len(defaultValues) {
			values[field.Name] = defaultValues[i]
		} else {
			values[field.Name] = ""
		}
	}

	model := FormModel{
		fields: fields,
		cursor: 0,
		values: values,
		errors: make(map[string]string),
	}

	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run form: %w", err)
	}

	formModel := finalModel.(FormModel)
	if formModel.quitting {
		return nil, fmt.Errorf("user cancelled")
	}

	return formModel.values, nil
}

// Helper function to check if string is numeric
func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
