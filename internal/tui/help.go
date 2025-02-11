package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
)

func (m model) renderHelpView() string {
	var (
		firstRow  []key.Binding
		secondRow []key.Binding
		thirdRow  []key.Binding
	)

	switch m.view {
	case infoMessageView:
		if m.msg == "Successfully collected files" {
			firstRow = []key.Binding{
				m.keymap.viewCollectedFiles,
			}
			secondRow = []key.Binding{
				m.keymap.back,
				m.keymap.quit,
			}
		} else {
			firstRow = []key.Binding{
				m.keymap.back,
				m.keymap.quit,
			}
		}
	case initialView:
		firstRow = []key.Binding{
			m.keymap.up,
			m.keymap.down,
			m.keymap.enter,
			m.keymap.quit,
		}
	case listCollectedFilesView:
		firstRow = []key.Binding{
			// m.keymap.expandDir,
			m.keymap.delete,
			m.keymap.collectFiles,
		}
		secondRow = []key.Binding{
			m.keymap.up,
			m.keymap.down,
			// m.keymap.search,
		}
		thirdRow = []key.Binding{
			m.keymap.back,
			m.keymap.quit,
		}
	case listPathsView:
		firstRow = []key.Binding{
			m.keymap.add,
			// m.keymap.edit,
			m.keymap.delete,
			m.keymap.collectFiles,
		}
		secondRow = []key.Binding{
			m.keymap.up,
			m.keymap.down,
			// m.keymap.search,
		}
		thirdRow = []key.Binding{
			m.keymap.back,
			m.keymap.quit,
		}
	case listIgnorePatternsView:
		firstRow = []key.Binding{
			m.keymap.add,
			// m.keymap.edit,
			m.keymap.delete,
		}
		secondRow = []key.Binding{
			m.keymap.up,
			m.keymap.down,
			// m.keymap.search,
		}
		thirdRow = []key.Binding{
			m.keymap.back,
			m.keymap.quit,
		}
	case addPathView:
		fallthrough
	case addIgnorePatternView:
		firstRow = []key.Binding{
			m.keymap.inputSubmit,
			m.keymap.inputCancel,
		}
	case removePathsView:
		fallthrough
	case removeIgnorePatternsView:
		firstRow = []key.Binding{
			m.keymap.selectionToggle,
			m.keymap.selectionCancel,
			m.keymap.delete,
		}
		secondRow = []key.Binding{
			m.keymap.up,
			m.keymap.down,
			// m.keymap.search,
		}
		thirdRow = []key.Binding{
			m.keymap.back,
			m.keymap.quit,
		}
	default:
		firstRow = []key.Binding{
			m.keymap.up,
			m.keymap.down,
			m.keymap.enter,
		}
		secondRow = []key.Binding{
			m.keymap.back,
			m.keymap.quit,
		}
	}

	return fmt.Sprintf(
		"\n   %s\n   %s\n   %s\n",
		m.help.ShortHelpView(firstRow),
		m.help.ShortHelpView(secondRow),
		m.help.ShortHelpView(thirdRow),
	)
}
