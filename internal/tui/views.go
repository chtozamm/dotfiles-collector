package tui

import (
	"path/filepath"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

func buildTree(root *tree.Tree, files []fileops.File, depth int) *tree.Tree {
	for _, file := range files {
		childNode := tree.Root(filepath.Base(file.Path))

		// TODO: make tree interactive in TUI (allow movement, expanding directories)

		if len(file.Children) > 0 && depth > 1 {
			childNode = buildTree(childNode, file.Children, depth-1)
		}

		root = root.Child(childNode)
	}

	return root
}

func CollectedFilesTree(app *app.App, depth int) string {
	files, err := fileops.ListFiles(app.Destination)
	if err != nil {
		return ""
	}

	t := tree.Root(app.Destination)
	t = buildTree(t, files, depth)

	return t.String()
}
