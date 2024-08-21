package cmd

import (
	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

func setupListCmd(app *app.App, rootCmd *cobra.Command) {
	var listFiles = &cobra.Command{
		Use:   "list",
		Short: "Print the collected files and directories",
		Long:  `Print the collected files and directories.`,
		Run: func(cmd *cobra.Command, args []string) {
			app.ListCollectedFiles()
		},
	}

	rootCmd.AddCommand(listFiles)
}
