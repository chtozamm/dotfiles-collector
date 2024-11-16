package tui

import (
	"fmt"
	"log"
	"strings"

	"github.com/chtozamm/dotfiles-collector/internal/app"

	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	entryView sessionState = iota
	collectView
	collectedFilesView
	pathsView
	listPathsView
	addPathView
	removePathView
	ignorePatternsView
	addPatternView
	removePatternView
	listIgnorePatternsView
)

type model struct {
	app     *app.App
	state   sessionState
	choices map[sessionState][]string
	cursors map[sessionState]int
	err     error
}

func initialModel(app *app.App) model {
	choices := make(map[sessionState][]string)
	choices[entryView] = []string{
		"Collect files",
		"List collected files",
		"Manage paths to collect from",
		"Manage ignored patterns",
	}

	choices[pathsView] = []string{
		"List paths",
		"Add new path",
		"Remove path",
	}

	choices[ignorePatternsView] = []string{
		"List ignore patterns",
		"Add new pattern",
		"Remove pattern",
	}
	return model{
		state:   entryView,
		app:     app,
		choices: choices,
		cursors: make(map[sessionState]int),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.moveCursor(-1)
		case "down", "j":
			m.moveCursor(1)
		case "enter", " ":
			m.handleEnter()
		case "backspace":
			m.handleBackspace()
		}
	}
	return m, nil
}

func (m *model) moveCursor(direction int) {
	if direction < 0 && m.cursors[m.state] > 0 {
		m.cursors[m.state]--
	} else if direction > 0 && m.cursors[m.state] < len(m.choices[m.state])-1 {
		m.cursors[m.state]++
	}
}

func (m *model) handleEnter() {
	switch m.state {
	case entryView:
		m.handleEntryView()
	case pathsView:
		m.handlePathsView()
	case ignorePatternsView:
		m.handleIgnorePatternsView()
		// case removePathView:
		// 	m.err = m.app.RemoveCollectPath(int64(m.cursors[m.state]))
		// 	m.state = listPathsView
	}
}

func (m *model) handleEntryView() {
	switch m.cursors[m.state] {
	case 0:
		m.state = collectView
		m.err = m.app.CopyFiles()
	case 1:
		m.state = collectedFilesView
	case 2:
		m.state = pathsView
	case 3:
		m.state = ignorePatternsView
	}
}

func (m *model) handlePathsView() {
	switch m.cursors[m.state] {
	case 0:
		m.state = listPathsView
		// case 2:
		// 	m.state = removePathView
	}
}

func (m *model) handleIgnorePatternsView() {
	switch m.cursors[m.state] {
	case 0:
		m.state = listIgnorePatternsView
	}
}

func (m *model) handleBackspace() {
	switch m.state {
	case collectView, pathsView, collectedFilesView, ignorePatternsView:
		m.state = entryView
	case listPathsView:
		m.state = pathsView
	case listIgnorePatternsView:
		m.state = ignorePatternsView
	}
}

func (m model) View() string {
	s := "Dotfiles Collector\n\n"
	s += "At the moment, only \"collect\" and \"list\" functions are implemented in TUI.\n"
	s += "For CLI info, use: dotfiles-collect -h \n\n"
	switch m.state {
	case entryView:
		return m.renderMenu(s, entryView)
	case collectView:
		return m.renderCollectView(s)
	case collectedFilesView:
		return m.renderCollectedFilesView(s)
	case pathsView:
		return m.renderMenu(s, pathsView)
	case listPathsView:
		return m.renderListPathsView(s)
	case ignorePatternsView:
		return m.renderMenu(s, ignorePatternsView)
	case listIgnorePatternsView:
		return m.renderListIgnorePatternsView(s)
	}
	return "Unknown state!"
}
func (m model) renderMenu(s string, state sessionState) string {
	for i, choice := range m.choices[state] {
		cursor := " "
		if m.cursors[state] == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress backspace to go back or q to quit.\n"
	return s
}

func (m model) renderCollectView(s string) string {
	if m.err != nil {
		s += fmt.Sprintf("Failed to collect files: %v\n", m.err)
	} else {
		s += "Successfully collected the files.\n"
	}
	s += "\nPress backspace to go back or q to quit.\n"
	return s
}

func (m model) renderCollectedFilesView(s string) string {
	var sb strings.Builder
	sb.WriteString(s)
	sb.WriteString("Collected files:\n\n")
	sb.WriteString(CollectedFilesTree(m.app, 1))
	sb.WriteString("\n\nPress backspace to go back or q to quit.\n")
	return sb.String()
}

func (m model) renderListPathsView(s string) string {
	var sb strings.Builder
	sb.WriteString(s)
	sb.WriteString("Paths to collect:\n\n")
	for _, path := range m.app.GetCollectPaths() {
		sb.WriteString(path.Path)
		if path.ParentDir != "" {
			sb.WriteString(", parent: " + path.ParentDir)
		}
		sb.WriteString("\n")
	}
	sb.WriteString("\nPress backspace to go back or q to quit.\n")
	return sb.String()
}

func (m model) renderListIgnorePatternsView(s string) string {
	var sb strings.Builder
	sb.WriteString(s)
	sb.WriteString("Ignore patterns:\n\n")
	for _, pattern := range m.app.GetIgnorePatterns() {
		sb.WriteString(pattern + "\n")
	}
	sb.WriteString("\nPress backspace to go back or q to quit.\n")
	return sb.String()
}

func Execute(app *app.App) {
	p := tea.NewProgram(initialModel(app))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
