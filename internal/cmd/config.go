package cmd

import (
	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

func setupConfigCmd(app *app.App, rootCmd *cobra.Command) {
	configCmd := &cobra.Command{
		Use:   "config <list>",
		Short: "Manage configuration of the application",
		Long:  "Manage configuration of the application.",
	}

	listConfig := &cobra.Command{
		Use:   "list",
		Short: "List configuration settings",
		Long:  "List configuration settings.",
		Run: func(cmd *cobra.Command, args []string) {
			app.ListConfig()
		},
	}

	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(listConfig)
}
