package cmd

import (
	"fmt"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/chtozamm/dotfiles-collector/internal/tui"
	"github.com/spf13/cobra"
)

func setupListCmd(app *app.App, rootCmd *cobra.Command) {
	listFiles := &cobra.Command{
		Use:   "list",
		Short: "List collected files",
		Long:  `List collected files.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(tui.CollectedFilesTree(app, 1))
		},
	}

	rootCmd.AddCommand(listFiles)
}
