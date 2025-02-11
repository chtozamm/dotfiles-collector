package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

func (m model) renderMenuView() string {
	sb := strings.Builder{}
	for i, choice := range m.options[m.view] {
		cursor := " "
		if m.cursors[m.view] == i {
			cursor = cursorStyle.Render(">")
		}
		sb.WriteString(fmt.Sprintf("%s%s %s\n", m.marginLeft, cursor, choice))
	}
	return sb.String()
}

func (m model) renderAddPathView() string {
	return fmt.Sprintf(
		"%s  Add new path entry:\n\n  %s\n",
		m.marginLeft,
		m.textInput.View(),
	)
}

func (m model) renderAddIgnorePatternView() string {
	return fmt.Sprintf(
		"%s  Add new ignore pattern:\n\n  %s\n",
		m.marginLeft,
		m.textInput.View(),
	)
}

// func (m model) renderConfirmationPromptView() string {
// 	sb := strings.Builder{}
// 	sb.WriteString(fmt.Sprintf("  %v\n", m.msg))
// 	return sb.String()
// }

func (m model) renderInfoMessageView() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("  %s%v\n", m.marginLeft, m.msg))
	return sb.String()
}

func (m *model) renderCollectedFilesView() string {
	files, err := fileops.ListFiles(m.app.Destination)
	if err != nil {
		m.msg = fmt.Sprintf("%s  Failed to get collected files: %v\n", m.marginLeft, err)
		m.view = infoMessageView
		return m.msg
	}

	if len(files) == 0 {
		m.msg = fmt.Sprintf("%s  No files have been collected yet\n", m.marginLeft)
		m.view = infoMessageView
		return m.msg
	}

	filenames := make([]string, 0, len(files))
	for _, file := range files {
		filenames = append(filenames, file.Path)
	}
	m.options[m.view] = filenames

	sb := strings.Builder{}
	sb.WriteString(m.marginLeft)
	sb.WriteString("  Collected files:\n\n")
	for i, file := range files {
		cursor := " "
		if m.cursors[m.view] == i {
			cursor = cursorStyle.Render(">")
		}
		if file.IsDir {
			sb.WriteString(fmt.Sprintf("%s%s %s\n", m.marginLeft, cursor, dirStyle.Render(filepath.Base(file.Path))))
		} else {
			sb.WriteString(fmt.Sprintf("%s%s %s\n", m.marginLeft, cursor, filepath.Base(file.Path)))
		}
	}
	return sb.String()
}

func (m *model) renderListPathsView() string {
	paths, err := m.app.GetCollectPaths()
	if err != nil {
		m.msg = fmt.Sprintf("%s  Failed to get paths: %v\n", m.marginLeft, err)
		m.lastView = m.view
		m.view = infoMessageView
		return m.msg
	}

	pathNames := make([]string, 0, len(paths))
	for _, path := range paths {
		pathNames = append(pathNames, path.Path)
	}
	m.options[m.view] = pathNames

	sb := strings.Builder{}
	if len(pathNames) > 0 {
		sb.WriteString(fmt.Sprintf("%s  Paths to collect files from:\n\n", m.marginLeft))
	} else {
		m.msg = fmt.Sprintf("%s  You haven't added any paths to collect files from\n", m.marginLeft)
		m.lastView = managePathsView
		m.view = infoMessageView
		return m.msg
	}
	for i, path := range paths {
		cursor := " "
		if m.cursors[m.view] == i {
			cursor = cursorStyle.Render(">")
		}
		if path.Subdir != "" {
			sb.WriteString(fmt.Sprintf("%s%s %s ⟶  %s\n", m.marginLeft, cursor, entryStyle.Render(path.Path), path.Subdir))
		} else {
			sb.WriteString(fmt.Sprintf("%s%s %s\n", m.marginLeft, cursor, entryStyle.Render(path.Path)))
		}
	}
	return sb.String()
}

func (m *model) renderListIgnorePatternsView() string {
	patterns, err := m.app.GetIgnorePatterns()
	if err != nil {
		m.msg = fmt.Sprintf("%s  Failed to get ignore patterns: %v\n", m.marginLeft, err)
		m.lastView = m.view
		m.view = infoMessageView
		return m.msg
	}

	m.options[m.view] = patterns

	sb := strings.Builder{}
	if len(patterns) > 0 {
		sb.WriteString(fmt.Sprintf("%s  Patterns for collector to ignore:\n\n", m.marginLeft))
	} else {
		m.msg = fmt.Sprintf("%s  You haven't added any ignore patterns yet\n", m.marginLeft)
		m.lastView = manageIgnorePatternsView
		m.view = infoMessageView
		return m.msg
	}
	for i, pattern := range patterns {
		cursor := " "
		if m.cursors[m.view] == i {
			cursor = cursorStyle.Render(">")
		}
		sb.WriteString(fmt.Sprintf("%s%s %s\n", m.marginLeft, cursor, entryStyle.Render(pattern)))
	}
	return sb.String()
}

func (m *model) renderRemovePathsView() string {
	paths, err := m.app.GetCollectPaths()
	if err != nil {
		m.msg = fmt.Sprintf("%s  Failed to get paths: %v\n", m.marginLeft, err)
		m.lastView = m.view
		m.view = infoMessageView
		return m.msg
	}
	pathNames := make([]string, 0, len(paths))
	for _, path := range paths {
		pathNames = append(pathNames, path.Path)
	}
	m.options[m.view] = pathNames

	sb := strings.Builder{}
	if len(pathNames) > 0 {
		sb.WriteString(fmt.Sprintf("%s  Paths to collect files from:\n\n", m.marginLeft))
	} else {
		m.msg = fmt.Sprintf("%s  You haven't added any paths to collect files from\n", m.marginLeft)
		m.lastView = managePathsView
		m.view = infoMessageView
		return m.msg
	}
	for i, path := range paths {
		cursor := " "
		if m.cursors[m.view] == i {
			cursor = cursorStyle.Render(">")
		}
		lb := lipgloss.NewStyle().Foreground(lipgloss.Color("#424243")).Render("[")
		rb := lipgloss.NewStyle().Foreground(lipgloss.Color("#424243")).Render("]")
		checked := " "
		if m.choices[m.view][path.Path] {
			checked = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")).Render("x")
		}
		if path.Subdir != "" {
			sb.WriteString(fmt.Sprintf("%s%s %s%s%s %s ⟶  %s\n", m.marginLeft, cursor, lb, checked, rb, entryStyle.Render(path.Path), path.Subdir))
		} else {
			sb.WriteString(fmt.Sprintf("%s%s %s%s%s %s\n", m.marginLeft, cursor, lb, checked, rb, entryStyle.Render(path.Path)))
		}
	}
	return sb.String()
}

func (m *model) renderRemoveIgnorePatternsView() string {
	patterns, err := m.app.GetIgnorePatterns()
	if err != nil {
		m.msg = fmt.Sprintf("%s  Failed to get ignore patterns: %v\n", m.marginLeft, err)
		m.lastView = m.view
		m.view = infoMessageView
		return m.msg
	}
	m.options[m.view] = patterns
	sb := strings.Builder{}
	if len(patterns) > 0 {
		sb.WriteString(fmt.Sprintf("%s  Patterns for collector to ignore:\n\n", m.marginLeft))
	} else {
		m.msg = fmt.Sprintf("%s  You haven't added any ignore patterns yet\n", m.marginLeft)
		m.lastView = manageIgnorePatternsView
		m.view = infoMessageView
		return m.msg
	}
	for i, pattern := range patterns {
		cursor := " "
		if m.cursors[m.view] == i {
			cursor = cursorStyle.Render(">")
		}
		lb := lipgloss.NewStyle().Foreground(lipgloss.Color("#424243")).Render("[")
		rb := lipgloss.NewStyle().Foreground(lipgloss.Color("#424243")).Render("]")
		checked := " "
		if m.choices[m.view][pattern] {
			checked = lipgloss.NewStyle().Foreground(lipgloss.Color("#f38ba8")).Render("x")
		}
		sb.WriteString(fmt.Sprintf("%s%s %s%s%s %s\n", m.marginLeft, cursor, lb, checked, rb, entryStyle.Render(pattern)))
	}
	return sb.String()
}
