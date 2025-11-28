package tui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ProgressModel represents a progress bar state
type ProgressModel struct {
	steps    []string
	current  int
	total    int
	percent  float64
	running  bool
	quitting bool
}

// ProgressOptions contains configuration for progress
type ProgressOptions struct {
	Steps   []string
	Current int
}

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

// ProgressTickMsg is sent to update progress
type ProgressTickMsg struct{}

// Init initializes the progress model
func (m ProgressModel) Init() tea.Cmd {
	if m.running {
		return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
			return ProgressTickMsg{}
		})
	}
	return nil
}

// Update handles messages and updates the model
func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit
		}

	case ProgressTickMsg:
		if m.running && m.current < m.total-1 {
			m.current++
			m.percent = float64(m.current+1) / float64(m.total) * 100

			if m.current >= m.total-1 {
				m.running = false
				return m, tea.Quit
			}

			return m, tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
				return ProgressTickMsg{}
			})
		}
	}

	return m, nil
}

// View renders the progress bar
func (m ProgressModel) View() string {
	if m.quitting {
		return ""
	}

	var content strings.Builder

	// Title
	content.WriteString(stepStyle.Render("Progress"))
	content.WriteString("\n\n")

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
	if m.running {
		content.WriteString("\n\n")
		content.WriteString(infoStyle.Render("Press Esc to cancel"))
	} else {
		content.WriteString("\n\n")
		content.WriteString(completeStepStyle.Render("✓ Complete!"))
	}

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#007DCC")).
		Padding(1, 2).
		Render(content.String())
}

// ShowProgress displays a progress bar
func ShowProgress(steps []string, current int) error {
	if len(steps) == 0 {
		return fmt.Errorf("no steps provided")
	}

	if current < 0 {
		current = 0
	} else if current >= len(steps) {
		current = len(steps) - 1
	}

	model := ProgressModel{
		steps:   steps,
		current: current,
		total:   len(steps),
		percent: float64(current+1) / float64(len(steps)) * 100,
		running: current < len(steps)-1,
	}

	p := tea.NewProgram(model)
	_, err := p.Run()
	if err != nil {
		return fmt.Errorf("failed to run progress: %w", err)
	}

	return nil
}
