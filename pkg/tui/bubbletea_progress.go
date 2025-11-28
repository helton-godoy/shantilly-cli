package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BubbleTeaProgressModel represents a progress indicator state
type BubbleTeaProgressModel struct {
	steps    []string
	current  int
	percent  float64
	running  bool
	quitting bool
}

// BubbleTeaProgressOptions contains configuration for progress
type BubbleTeaProgressOptions struct {
	Steps   []string
	Current int
}

// ProgressTickMsg is sent to update progress
type ProgressTickMsg struct{}

// Styles for progress
var (
	progressBarStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#E9ECEF")).
				Foreground(lipgloss.Color("#007DCC"))

	completeBarStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#007DCC")).
				Foreground(lipgloss.Color("white"))

	stepStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#007DCC"))

	completeStepStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#28A745"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C757D"))
)

// Init initializes the progress model
func (m BubbleTeaProgressModel) Init() tea.Cmd {
	// Don't auto-start, let user control
	return nil
}

// Update handles messages and updates the model
func (m BubbleTeaProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc, tea.KeyEnter, tea.KeySpace:
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

// View renders the progress indicator
func (m BubbleTeaProgressModel) View() string {
	if m.quitting {
		return ""
	}

	var content strings.Builder

	// Progress bar
	barWidth := 40
	filled := int(float64(barWidth) * m.percent / 100)
	empty := barWidth - filled

	progressBar := completeBarStyle.Render(strings.Repeat(" ", filled)) +
		progressBarStyle.Render(strings.Repeat(" ", empty))

	content.WriteString(progressBar)
	content.WriteString(fmt.Sprintf(" %.1f%%", m.percent))
	content.WriteString("\n\n")

	// Steps
	content.WriteString(infoStyle.Render("Steps:"))
	content.WriteString("\n")

	for i, step := range m.steps {
		prefix := "○"
		style := infoStyle

		if i < m.current {
			prefix = "✓"
			style = completeStepStyle
		} else if i == m.current {
			prefix = "◉"
			style = stepStyle
		}

		content.WriteString(fmt.Sprintf("%s %s", prefix, style.Render(step)))
		content.WriteString("\n")
	}

	// Current step info
	if m.current < len(m.steps) {
		content.WriteString("\n")
		content.WriteString(infoStyle.Render(fmt.Sprintf("Current: %s", m.steps[m.current])))
	}

	// Instructions
	content.WriteString("\n\n")
	content.WriteString(infoStyle.Render("Press Enter/Space to close"))

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#007DCC")).
		Padding(1, 2).
		Render(content.String())
}

// BubbleTeaShowProgress displays a BubbleTea progress indicator
func BubbleTeaShowProgress(steps []string, current int) error {
	if len(steps) == 0 {
		return fmt.Errorf("no steps provided")
	}

	// Validate current step
	if current < 0 {
		current = 0
	} else if current >= len(steps) {
		current = len(steps) - 1
	}

	percent := float64(current+1) / float64(len(steps)) * 100

	model := BubbleTeaProgressModel{
		steps:   steps,
		current: current,
		percent: percent,
		running: false, // Static display
	}

	p := tea.NewProgram(model)
	_, err := p.Run()
	if err != nil {
		return fmt.Errorf("failed to run progress: %w", err)
	}

	return nil
}
