package cmd

import (
	"fmt"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

func setupIgnoreCmd(app *app.App, rootCmd *cobra.Command) {
	var ignoreCmd = &cobra.Command{
		Use:   "ignore <add|list|remove>",
		Short: "Manage ignore patterns",
		Long:  "List, add or remove regexp patterns for the collector to ignore.",
	}

	var addPath = &cobra.Command{
		Use:   "add <pattern>",
		Short: "Add ignore pattern",
		Long:  `Add ignore pattern for the collector to ignore.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(`Usage:
  dotfiles-collector ignore add <pattern>`)
				return
			}
			app.AddIgnorePattern(args[0])
		},
	}

	var listIgnorePatterns = &cobra.Command{
		Use:   "list",
		Short: "List ignore patterns added to the collector",
		Long:  "List ignore patterns added to the collector.",
		Run: func(cmd *cobra.Command, args []string) {
			app.ListIgnorePatterns()
		}}

	rootCmd.AddCommand(ignoreCmd)
	ignoreCmd.AddCommand(addPath)
	ignoreCmd.AddCommand(listIgnorePatterns)
}
