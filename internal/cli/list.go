package cli

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
	"github.com/spf13/cobra"
)

func setupListCmd(app *app.Application, rootCmd *cobra.Command) {
	listFiles := &cobra.Command{
		Use:   "list",
		Short: "List collected files",
		Long:  `List collected files.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(collectedFilesTree(app, 1))
		},
	}

	rootCmd.AddCommand(listFiles)
}

func buildTree(root *tree.Tree, files []fileops.File, depth int) *tree.Tree {
	for _, file := range files {
		childNode := tree.Root(filepath.Base(file.Path))

		if len(file.Children) > 0 && depth > 1 {
			childNode = buildTree(childNode, file.Children, depth-1)
		}

		root = root.Child(childNode)
	}

	return root
}

func collectedFilesTree(app *app.Application, depth int) string {
	files, err := fileops.ListFiles(app.Destination)
	if err != nil {
		return ""
	}

	t := tree.Root(app.Destination)
	t = buildTree(t, files, depth)

	return t.String()
}
