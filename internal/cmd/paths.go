package cmd

import (
	"fmt"
	"strconv"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

func setupPathsCmd(app *app.App, rootCmd *cobra.Command) {
	var pathsCmd = &cobra.Command{
		Use:   "paths <add|list|remove>",
		Short: "Manage source paths",
		Long:  "List, add or remove source path from the collector.",
	}

	var addPath = &cobra.Command{
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
			app.AddPath(args[0], parentDir)
		},
	}

	var removePath = &cobra.Command{
		Use:   "remove <path>",
		Short: "Remove path from the collector",
		Long:  `Remove path from the collector.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(`Usage:
  dotfiles-collector paths remove <path>`)
				return
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid path ID:", id)
			}
			app.RemovePath(int64(id))
		},
	}

	var listPaths = &cobra.Command{
		Use:   "list",
		Short: "List source paths added to the collector",
		Long:  "List source paths added to the collector.",
		Run: func(cmd *cobra.Command, args []string) {
			app.ListPaths()
		}}

	rootCmd.AddCommand(pathsCmd)
	pathsCmd.AddCommand(addPath)
	pathsCmd.AddCommand(removePath)
	pathsCmd.AddCommand(listPaths)
}
