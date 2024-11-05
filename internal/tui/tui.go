package tui

import (
	"fmt"
	"log"
	"strings"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"

	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	entryView sessionState = iota
	collectView
	collectedFilesView
	pathsView
	listPathsView
	ignorePatternsView
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
			if m.cursors[m.state] > 0 {
				m.cursors[m.state]--
			}

		case "down", "j":
			if m.cursors[m.state] < len(m.choices[m.state])-1 {
				m.cursors[m.state]++
			}

		case "enter", " ":
			switch m.state {
			case entryView:
				switch m.cursors[m.state] {
				case 1:
					m.state = collectedFilesView
				case 2:
					m.state = pathsView
				case 3:
					m.state = ignorePatternsView
				}
			case pathsView:
				switch m.cursors[m.state] {
				case 0:
					m.state = listPathsView
				}
			case ignorePatternsView:
				switch m.cursors[m.state] {
				case 0:
					m.state = listIgnorePatternsView
				}
			}
		case "backspace":
			switch m.state {
			case pathsView, collectedFilesView, ignorePatternsView:
				m.state = entryView
			case listPathsView:
				m.state = pathsView
			case listIgnorePatternsView:
				m.state = ignorePatternsView
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Dotfiles Collector\n\n"
	switch m.state {
	case entryView, pathsView, ignorePatternsView:

		for i, choice := range m.choices[m.state] {
			cursor := " "
			if m.cursors[m.state] == i {
				cursor = ">"
			}

			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}

		// The footer
		if m.state == entryView {
			s += "\nPress q to quit.\n"
		} else {
			s += "\nPress backspace to go back or q to quit.\n"
		}

		// Send the UI for rendering
		return s

	case collectedFilesView:
		files, err := fileops.ListFiles(m.app.Destination, 1)
		if err != nil {
			return err.Error()
		}
		var sb strings.Builder
		sb.WriteString(s)
		sb.WriteString("Collected files:\n\n")
		for _, file := range files {
			sb.WriteString(file + "\n")
		}
		sb.WriteString("\nPress backspace to go back or q to quit.\n")
		return sb.String()

	case listPathsView:
		var sb strings.Builder
		sb.WriteString(s)
		sb.WriteString("Paths to collect:\n\n")
		for _, path := range m.app.SourcePaths {
			sb.WriteString(path.Path + "\n")
		}
		sb.WriteString("\nPress backspace to go back or q to quit.\n")
		return sb.String()

	case listIgnorePatternsView:
		var sb strings.Builder
		sb.WriteString(s)
		sb.WriteString("Ignore patterns:\n\n")
		for pattern := range m.app.IgnorePatterns {
			sb.WriteString(pattern + "\n")
		}
		sb.WriteString("\nPress backspace to go back or q to quit.\n")
		return sb.String()
	}
	return "Hello there!"
}

func Execute(app *app.App) {
	p := tea.NewProgram(initialModel(app))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
