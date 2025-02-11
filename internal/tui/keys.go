package tui

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	up    key.Binding
	down  key.Binding
	enter key.Binding
	back  key.Binding
	// search          key.Binding
	viewCollectedFiles key.Binding
	collectFiles       key.Binding
	selectionToggle    key.Binding
	selectionCancel    key.Binding
	inputSubmit        key.Binding
	inputCancel        key.Binding
	quit               key.Binding
	add                key.Binding
	edit               key.Binding
	delete             key.Binding
	expandDir          key.Binding
}

var keymaps = keymap{
	up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	enter: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter/space", "select"),
	),
	inputSubmit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm"),
	),
	inputCancel: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("esc", "cancel"),
	),
	back: key.NewBinding(
		key.WithKeys("backspace", "esc"),
		key.WithHelp("backspace", "previous screen"),
	),
	// search: key.NewBinding(
	// 	key.WithKeys("/"),
	// 	key.WithHelp("/", "search"),
	// ),
	viewCollectedFiles: key.NewBinding(
		key.WithKeys("l"),
		key.WithHelp("l", "view collected files"),
	),
	collectFiles: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "collect files"),
	),
	selectionToggle: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter/space", "select/deselect"),
	),
	selectionCancel: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "deselect all"),
	),
	quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	add: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add"),
	),
	edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
	),
	delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
	expandDir: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter/space", "expand directory"),
	),
}
