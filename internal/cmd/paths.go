package cmd

import (
	"fmt"
	"strings"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

func setupPathsCmd(app *app.App, rootCmd *cobra.Command) {
	pathsCmd := &cobra.Command{
		Use:   "paths <add|list|remove>",
		Short: "Manage source paths",
		Long:  "List, add or remove source paths from the collector.",
	}

	addPath := &cobra.Command{
		Use:   "add <path>",
		Short: "Add path to the collector",
		Long:  `Add path to the collector.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(`Usage:
  dotfiles-collector paths add <path>`)
				return
			}
			var parentDir string
			if len(args) == 2 {
				parentDir = args[1]
			}
			err := app.AddCollectPath(args[0], parentDir)
			if err != nil {
				fmt.Printf("Failed to add path: %s\n", err)
				return
			}
		},
	}

	removePath := &cobra.Command{
		Use:   "remove <path>",
		Short: "Remove path from the collector",
		Long:  `Remove path from the collector.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(`Usage:
  dotfiles-collector paths remove <path>`)
				return
			}
			err := app.RemoveCollectPath(args[0])
			if err != nil {
				fmt.Printf("Failed to remove path: %s\n", err)
			}
		},
	}

	listPaths := &cobra.Command{
		Use:   "list",
		Short: "List source paths added to the collector",
		Long:  "List source paths added to the collector.",
		Run: func(cmd *cobra.Command, args []string) {
			var sb strings.Builder
			for _, path := range app.GetCollectPaths() {
				sb.WriteString(path.Path)
				if path.ParentDir != "" {
					sb.WriteString(", parent: " + path.ParentDir)
				}
				sb.WriteString("\n")
			}
			fmt.Print(sb.String())
		},
	}

	rootCmd.AddCommand(pathsCmd)
	pathsCmd.AddCommand(addPath)
	pathsCmd.AddCommand(removePath)
	pathsCmd.AddCommand(listPaths)
}
