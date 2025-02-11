package tui

import (
	"fmt"
	"strings"

	"github.com/chtozamm/dotfiles-collector/internal/app"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewState int

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff944e"))
	dirStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#f5e0dc")).Bold(true)
	entryStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#f5e0dc"))
	titleStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff944e")).Border(lipgloss.Border{Bottom: "─"}).MarginLeft(2)
)

const (
	initialView viewState = iota
	listCollectedFilesView

	infoMessageView

	// promptConfirmAdd
	// promptConfirmEdit
	// promptConfirmDelete

	managePathsView
	listPathsView
	addPathView
	// addPathParentView
	removePathsView

	manageIgnorePatternsView
	listIgnorePatternsView
	addIgnorePatternView
	removeIgnorePatternsView
)

type model struct {
	app        *app.Application
	view       viewState
	lastView   viewState
	options    map[viewState][]string
	choices    map[viewState]map[string]bool
	cursors    map[viewState]int
	textInput  textinput.Model
	msg        string
	keymap     keymap
	help       help.Model
	marginLeft string
}

func initialModel(app *app.Application) model {
	options := make(map[viewState][]string)
	choices := make(map[viewState]map[string]bool)

	options[initialView] = []string{
		"Collect files",
		"List collected files",
		"Manage paths to collect from",
		"Manage ignored patterns",
	}

	options[managePathsView] = []string{
		"List paths",
		"Add new path",
		"Remove paths",
	}

	options[manageIgnorePatternsView] = []string{
		"List ignore patterns",
		"Add new pattern",
		"Remove patterns",
	}

	marginLeft := " "

	textInput := textinput.New()
	textInput.Prompt = cursorStyle.Render(fmt.Sprintf("%s> ", marginLeft))
	textInput.Placeholder = "Waiting for your input..."
	textInput.Focus()
	textInput.Width = 44

	m := model{
		app:        app,
		view:       initialView,
		lastView:   initialView,
		options:    options,
		choices:    choices,
		cursors:    make(map[viewState]int),
		keymap:     keymaps,
		help:       help.New(),
		textInput:  textInput,
		marginLeft: marginLeft,
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.SetWindowTitle("Dotfiles Collector"), tea.Cmd(tea.ClearScreen))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch m.view {
		default:
		case addPathView:
			fallthrough
		case addIgnorePatternView:
			return handleInput(&m, msg)
		case infoMessageView:
			if key.Matches(msg, m.keymap.viewCollectedFiles) {
				if m.msg != "Successfully collected files" {
					break
				}
				m.lastView = m.view
				m.view = listCollectedFilesView
			}
		case listCollectedFilesView:
			fallthrough
		case listPathsView:
			if key.Matches(msg, m.keymap.collectFiles) {
				m.handleCollectFiles()
			}
		}

		switch {
		case key.Matches(msg, m.keymap.quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.up):
			m.moveCursor(-1)
		case key.Matches(msg, m.keymap.down):
			m.moveCursor(1)
		case key.Matches(msg, m.keymap.enter):
			m.handleEnter()
		case key.Matches(msg, m.keymap.add):
			m.handleAdd()
		case key.Matches(msg, m.keymap.delete):
			m.handleDelete()
		case key.Matches(msg, m.keymap.back):
			m.handleBackspace()
		case key.Matches(msg, m.keymap.selectionToggle):
			m.handleSelectionToggle()
		case key.Matches(msg, m.keymap.selectionCancel):
			m.handleSelectionCancel()
		case key.Matches(msg, m.keymap.viewCollectedFiles):
		}
	}

	return m, nil
}

func (m model) View() string {
	// Render header
	sb := strings.Builder{}
	sb.WriteString(titleStyle.Render("Dotfiles Collector"))
	sb.WriteString("\n\n")

	// Render content
	switch m.view {
	case infoMessageView:
		sb.WriteString(m.renderInfoMessageView())
	case listCollectedFilesView:
		sb.WriteString(m.renderCollectedFilesView())
	case listPathsView:
		sb.WriteString(m.renderListPathsView())
	case addPathView:
		sb.WriteString(m.renderAddPathView())
	case listIgnorePatternsView:
		sb.WriteString(m.renderListIgnorePatternsView())
	case addIgnorePatternView:
		sb.WriteString(m.renderAddIgnorePatternView())
	case removePathsView:
		sb.WriteString(m.renderRemovePathsView())
	case removeIgnorePatternsView:
		sb.WriteString(m.renderRemoveIgnorePatternsView())
	default:
		sb.WriteString(m.renderMenuView())
	}

	// Render footer
	sb.WriteString(m.renderHelpView())
	return sb.String()
}

// Execute starts a terminal user interface.
func Execute(app *app.Application) error {
	p := tea.NewProgram(initialModel(app))
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}
