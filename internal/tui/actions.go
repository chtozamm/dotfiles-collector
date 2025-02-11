package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

func (m *model) moveCursor(direction int) {
	if direction < 0 && m.cursors[m.view] > 0 {
		m.cursors[m.view]--
	} else if direction > 0 && m.cursors[m.view] < len(m.options[m.view])-1 {
		m.cursors[m.view]++
	}
}

func (m *model) handleCollectFiles() {
	m.lastView = m.view
	err := m.app.CopyFiles()
	if err != nil {
		if err.Error() == "no paths found in database" {
			m.msg = "No paths found to collect files from"
		} else {
			m.msg = fmt.Sprintf("Failed to collect files: %v", err)
		}
	} else {
		m.msg = "Successfully collected files"
	}
	m.view = infoMessageView
}

func (m *model) handleEnter() {
	switch m.view {
	case initialView:
		m.handleInitialView()
	case managePathsView:
		m.handlePathsView()
	case manageIgnorePatternsView:
		m.handleIgnorePatternsView()
	case removePathsView:
		fallthrough
	case removeIgnorePatternsView:
		m.handleSelectionToggle()
	}
}

func (m *model) handleDeletePath() {
	if len(m.options[m.view]) == 0 {
		return
	}

	err := m.app.RemoveCollectPath(m.options[m.view][m.cursors[m.view]])
	if err != nil {
		m.msg = fmt.Sprintf("Failed to get remove path: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}

	paths, err := m.app.GetCollectPaths()
	if err != nil {
		m.msg = fmt.Sprintf("Failed to get paths: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}
	pathNames := make([]string, 0, len(paths))
	for _, path := range paths {
		pathNames = append(pathNames, path.Path)
	}
	m.options[m.view] = pathNames
}

func (m *model) handleDeletePaths() {
	if len(m.options[m.view]) == 0 {
		return
	}

	pathsToDelete := []string{}

	for pathname, marked := range m.choices[m.view] {
		if !marked {
			continue
		}
		pathsToDelete = append(pathsToDelete, pathname)
	}

	err := m.app.RemoveCollectPaths(pathsToDelete)
	if err != nil {
		m.msg = fmt.Sprintf("Failed to remove paths: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}

	paths, err := m.app.GetCollectPaths()
	if err != nil {
		m.msg = fmt.Sprintf("Failed to get collect paths: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}
	pathNames := make([]string, 0, len(paths))
	for _, path := range paths {
		pathNames = append(pathNames, path.Path)
	}
	m.options[m.view] = pathNames
	m.choices[m.view] = map[string]bool{}
}

func (m *model) handleDelete() {
	switch m.view {
	case listCollectedFilesView:
		m.handleDeleteCollectedFile()
	case listPathsView:
		m.handleDeletePath()
	case removePathsView:
		m.handleDeletePaths()
	case listIgnorePatternsView:
		m.handleDeleteIgnorePattern()
	case removeIgnorePatternsView:
		m.handleDeleteIgnorePatterns()
	}

	if len(m.options[m.view]) == 0 {
		m.cursors[m.view] = 0
	} else if m.cursors[m.view] >= len(m.options[m.view]) {
		m.cursors[m.view] = len(m.options[m.view]) - 1
	}
}

func (m *model) handleDeleteCollectedFile() {
	if len(m.options[m.view]) == 0 {
		return
	}

	// m.err = nil
	// m.msg = fmt.Sprintf("Are you sure you want to delete %q?", filepath.Base(m.options[listCollectedFilesView][m.cursors[listCollectedFilesView]]))
	// m.view = infoMessageView

	err := fileops.Delete(m.options[listCollectedFilesView][m.cursors[listCollectedFilesView]])
	if err != nil {
		m.msg = fmt.Sprintf("Failed to delete file: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}

	files, err := fileops.ListFiles(m.app.Destination)
	if err != nil {
		m.msg = fmt.Sprintf("Failed to get collected files: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}
	filenames := make([]string, 0, len(files))
	for _, file := range files {
		filenames = append(filenames, file.Path)
	}
	m.options[m.view] = filenames
}

func (m *model) handleAdd() {
	switch m.view {
	case listPathsView:
		m.lastView = m.view
		m.view = addPathView
	case listIgnorePatternsView:
		m.lastView = m.view
		m.view = addIgnorePatternView
	}
}

func (m *model) handleBackspace() {
	m.cursors[m.view] = 0 // Reset cursor on the current view
	switch m.view {
	case infoMessageView:
		m.msg = ""
		m.view = m.lastView
	case listPathsView:
		m.view = managePathsView
	// case addPathView:
	// 	m.textInput.Reset()
	// 	m.view = managePathsView
	// case addIgnorePatternView:
	// 	m.textInput.Reset()
	// 	m.view = manageIgnorePatternsView
	case listIgnorePatternsView:
		m.view = manageIgnorePatternsView
	case removePathsView:
		m.choices[m.view] = map[string]bool{}
		m.view = managePathsView
	case removeIgnorePatternsView:
		m.choices[m.view] = map[string]bool{}
		m.view = manageIgnorePatternsView
	default:
		m.view = initialView
	}
}

func (m *model) handleSelectionCancel() {
	switch m.view {
	case removePathsView:
		fallthrough
	case removeIgnorePatternsView:
		m.choices[m.view] = map[string]bool{}
	}
}

func (m *model) handleSelectionToggle() {
	if m.choices[m.view] == nil {
		m.choices[m.view] = make(map[string]bool)
	}
	m.choices[m.view][m.options[m.view][m.cursors[m.view]]] = !m.choices[m.view][m.options[m.view][m.cursors[m.view]]]
}

func (m *model) handleInitialView() {
	switch m.cursors[m.view] {
	case 0:
		m.handleCollectFiles()
	case 1:
		m.view = listCollectedFilesView
	case 2:
		m.view = managePathsView
	case 3:
		m.view = manageIgnorePatternsView
	}
}

func (m *model) handlePathsView() {
	m.lastView = managePathsView
	switch m.cursors[m.view] {
	case 0:
		m.view = listPathsView
	case 1:
		m.view = addPathView
	case 2:
		m.view = removePathsView
	}
}

func (m *model) handleIgnorePatternsView() {
	m.lastView = manageIgnorePatternsView
	switch m.cursors[m.view] {
	case 0:
		m.view = listIgnorePatternsView
	case 1:
		m.view = addIgnorePatternView
	case 2:
		m.view = removeIgnorePatternsView
	}
}

func (m *model) handleDeleteIgnorePattern() {
	if len(m.options[m.view]) == 0 {
		return
	}

	err := m.app.RemoveIgnorePattern(m.options[m.view][m.cursors[m.view]])
	if err != nil {
		m.msg = fmt.Sprintf("Failed to remove ignore pattern: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}

	patterns, err := m.app.GetIgnorePatterns()
	if err != nil {
		m.msg = fmt.Sprintf("Failed to get ignore patterns: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}
	m.options[m.view] = patterns
}

func (m *model) handleDeleteIgnorePatterns() {
	if len(m.options[m.view]) == 0 {
		return
	}

	patternsToDelete := []string{}

	for pattern, marked := range m.choices[m.view] {
		if !marked {
			continue
		}
		patternsToDelete = append(patternsToDelete, pattern)
	}

	err := m.app.RemoveIgnorePatterns(patternsToDelete)
	if err != nil {
		m.msg = fmt.Sprintf("Failed to remove ignore patterns: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}

	patterns, err := m.app.GetIgnorePatterns()
	if err != nil {
		m.msg = fmt.Sprintf("Failed to get ignore patterns: %v", err)
		m.lastView = m.view
		m.view = infoMessageView
		return
	}
	m.options[m.view] = patterns
	m.choices[m.view] = map[string]bool{}
}

func handleInput(m *model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if key.Matches(msg, m.keymap.inputCancel) {
		m.textInput.Reset()
		m.view = m.lastView
		return m, nil
	}

	if key.Matches(msg, m.keymap.inputSubmit) {
		switch m.view {
		case addPathView:
			m.lastView = managePathsView
			err := m.app.AddCollectPath(m.textInput.Value(), "")
			if err != nil {
				m.textInput.Reset()
				m.msg = fmt.Sprintf("Failed to add path: %v", err)
				m.view = infoMessageView
				return m, nil
			}

			m.textInput.Reset()

			files, err := m.app.GetCollectPaths()
			if err != nil {
				m.textInput.Reset()
				m.msg = fmt.Sprintf("Failed to get collect paths: %v", err)
				m.view = infoMessageView
				return m, nil
			}

			filenames := make([]string, 0, len(files))
			for _, file := range files {
				filenames = append(filenames, file.Path)
			}

			m.options[listPathsView] = filenames
			m.view = listPathsView

			return m, nil
		case addIgnorePatternView:
			m.lastView = manageIgnorePatternsView
			err := m.app.AddIgnorePattern(m.textInput.Value())
			if err != nil {
				m.textInput.Reset()
				m.msg = fmt.Sprintf("Failed to add ignore pattern: %v", err)
				m.view = infoMessageView
				return m, nil
			}

			m.textInput.Reset()

			patterns, err := m.app.GetIgnorePatterns()
			if err != nil {
				m.textInput.Reset()
				m.msg = fmt.Sprintf("Failed to get ignore patterns: %v", err)
				m.view = infoMessageView
				return m, nil
			}

			m.options[listIgnorePatternsView] = patterns
			m.view = listIgnorePatternsView

			return m, nil
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}
