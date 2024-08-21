package cmd

import (
	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

func setupCollectCmd(app *app.App, rootCmd *cobra.Command) {
	var collectCmd = &cobra.Command{
		Use:   "collect",
		Short: "Collect files specified in source paths and copy to the destination",
		Long:  "Collect files specified in source paths and copy to the destination.",
		Run: func(cmd *cobra.Command, args []string) {
			app.CopyFiles()
		},
	}

	rootCmd.AddCommand(collectCmd)
}
