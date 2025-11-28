package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// FileInfo represents file information for selection
type FileInfo struct {
	Name    string
	Path    string
	IsDir   bool
	Size    int64
	ModTime string
}

// SelectorModel represents a file selector state
type SelectorModel struct {
	path     string
	files    []FileInfo
	cursor   int
	selected []string
	filter   string
	multi    bool
	quitting bool
	loading  bool
}

// SelectorOptions contains configuration for file selector
type SelectorOptions struct {
	Path   string
	Filter string
	Multi  bool
}

var (
	selectorTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#007DCC"))

	dirStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#007DCC"))

	fileStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#212529"))

	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("black")).
			Background(lipgloss.Color("#007DCC"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C757D"))

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#007DCC"))
)

// LoadFilesMsg is sent when files are loaded
type LoadFilesMsg struct {
	files []FileInfo
	err   error
}

// Init initializes the selector model
func (m SelectorModel) Init() tea.Cmd {
	return m.loadFiles()
}

// Update handles messages and updates the model
func (m SelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit

		case tea.KeyEnter:
			if len(m.files) > 0 && m.cursor >= 0 && m.cursor < len(m.files) {
				selectedFile := m.files[m.cursor]
				if selectedFile.IsDir {
					// Navigate into directory
					m.path = selectedFile.Path
					m.cursor = 0
					return m, m.loadFiles()
				} else {
					// Select file
					if m.multi {
						found := false
						for i, sel := range m.selected {
							if sel == selectedFile.Path {
								// Remove from selection
								m.selected = append(m.selected[:i], m.selected[i+1:]...)
								found = true
								break
							}
						}
						if !found {
							// Add to selection
							m.selected = append(m.selected, selectedFile.Path)
						}
					} else {
						// Single selection - return immediately
						m.selected = []string{selectedFile.Path}
						return m, tea.Quit
					}
				}
			}

		case tea.KeyUp, tea.KeyK:
			if m.cursor > 0 {
				m.cursor--
			}

		case tea.KeyDown, tea.KeyJ:
			if m.cursor < len(m.files)-1 {
				m.cursor++
			}

		case tea.KeyLeft, tea.KeyH:
			// Go to parent directory
			parent := filepath.Dir(m.path)
			if parent != m.path {
				m.path = parent
				m.cursor = 0
				return m, m.loadFiles()
			}

		case tea.KeySpace:
			if m.multi && len(m.files) > 0 && m.cursor >= 0 && m.cursor < len(m.files) {
				selectedFile := m.files[m.cursor]
				if !selectedFile.IsDir {
					found := false
					for i, sel := range m.selected {
						if sel == selectedFile.Path {
							// Remove from selection
							m.selected = append(m.selected[:i], m.selected[i+1:]...)
							found = true
							break
						}
					}
					if !found {
						// Add to selection
						m.selected = append(m.selected, selectedFile.Path)
					}
				}
			}

		case tea.KeyRight, tea.KeyL:
			if len(m.files) > 0 && m.cursor >= 0 && m.cursor < len(m.files) {
				selectedFile := m.files[m.cursor]
				if selectedFile.IsDir {
					// Navigate into directory
					m.path = selectedFile.Path
					m.cursor = 0
					return m, m.loadFiles()
				}
			}
		}

	case LoadFilesMsg:
		m.loading = false
		if msg.err != nil {
			m.files = []FileInfo{}
		} else {
			m.files = msg.files
		}
	}

	return m, nil
}

// View renders the file selector
func (m SelectorModel) View() string {
	if m.quitting {
		return ""
	}

	var content strings.Builder

	// Title
	content.WriteString(selectorTitleStyle.Render("File Selector"))
	content.WriteString("\n\n")

	// Current path
	content.WriteString(infoStyle.Render("Path: " + m.path))
	content.WriteString("\n\n")

	if m.loading {
		content.WriteString(infoStyle.Render("Loading..."))
	} else {
		// Files list
		maxItems := 15
		start := 0
		if m.cursor >= maxItems {
			start = m.cursor - maxItems + 1
		}
		end := start + maxItems
		if end > len(m.files) {
			end = len(m.files)
		}

		for i := start; i < end; i++ {
			file := m.files[i]

			// Selection indicator
			selected := ""
			if m.multi {
				for _, sel := range m.selected {
					if sel == file.Path {
						selected = "✓ "
						break
					}
				}
			}

			// Cursor indicator
			cursor := " "
			if i == m.cursor {
				cursor = "❯"
			}

			// File icon and styling
			icon := "📄 "
			style := fileStyle
			if file.IsDir {
				icon = "📁 "
				style = dirStyle
			}

			// File name
			line := fmt.Sprintf("%s %s%s%s", cursor, selected, icon, file.Name)

			// Apply styling
			if i == m.cursor {
				line = selectedStyle.Render(line)
			} else {
				line = style.Render(line)
			}

			content.WriteString(line)
			content.WriteString("\n")
		}

		// Selection info
		if m.multi && len(m.selected) > 0 {
			content.WriteString("\n")
			content.WriteString(infoStyle.Render(fmt.Sprintf("Selected: %d files", len(m.selected))))
		}
	}

	// Instructions
	content.WriteString("\n")
	content.WriteString(infoStyle.Render("↑↓ Navigate • Enter Select • Space Toggle (multi) • ← Parent • → Enter Dir • Esc Quit"))

	return borderStyle.Render(content.String())
}

// loadFiles loads files from the current path
func (m SelectorModel) loadFiles() tea.Cmd {
	m.loading = true
	return func() tea.Msg {
		entries, err := os.ReadDir(m.path)
		if err != nil {
			return LoadFilesMsg{err: err}
		}

		var files []FileInfo

		// Add parent directory entry
		parent := filepath.Dir(m.path)
		if parent != m.path {
			files = append(files, FileInfo{
				Name:  "..",
				Path:  parent,
				IsDir: true,
			})
		}

		// Read directory entries
		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				continue
			}

			filePath := filepath.Join(m.path, entry.Name())

			// Apply filter
			if m.filter != "" && !entry.IsDir() {
				matched, err := filepath.Match(m.filter, entry.Name())
				if err != nil || !matched {
					continue
				}
			}

			files = append(files, FileInfo{
				Name:    entry.Name(),
				Path:    filePath,
				IsDir:   entry.IsDir(),
				Size:    info.Size(),
				ModTime: info.ModTime().Format("2006-01-02"),
			})
		}

		// Sort files (directories first, then files)
		sort.Slice(files, func(i, j int) bool {
			if files[i].IsDir && !files[j].IsDir {
				return true
			}
			if !files[i].IsDir && files[j].IsDir {
				return false
			}
			return strings.ToLower(files[i].Name) < strings.ToLower(files[j].Name)
		})

		return LoadFilesMsg{files: files}
	}
}

// ShowFileSelector displays a file selector
func ShowFileSelector(path, filter string, multi bool) ([]string, error) {
	if path == "" {
		path = "."
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	model := SelectorModel{
		path:     absPath,
		filter:   filter,
		multi:    multi,
		selected: []string{},
		loading:  true,
	}

	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run file selector: %w", err)
	}

	selectorModel := finalModel.(SelectorModel)
	if selectorModel.quitting {
		return nil, fmt.Errorf("user cancelled")
	}

	return selectorModel.selected, nil
}
