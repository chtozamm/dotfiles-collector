package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

func setupIgnoreCmd(app *app.Application, rootCmd *cobra.Command) {
	ignoreCmd := &cobra.Command{
		Use:   "ignore <add|list|remove>",
		Short: "Manage ignore patterns",
		Long:  "List, add or remove ignore patterns as regular expressions for the collector to ignore if encountered.",
	}

	addPattern := &cobra.Command{
		Use:   "add <pattern>",
		Short: "Add ignore pattern",
		Long:  `Add ignore pattern.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(`Usage:
  dotfiles-collector ignore add <pattern>`)
				return
			}
			err := app.AddIgnorePattern(args[0])
			if err != nil {
				fmt.Printf("Failed to add ignore pattern: %s\n", err)
				os.Exit(1)
			}
		},
	}

	removePattern := &cobra.Command{
		Use:   "remove <pattern>",
		Short: "Remove ignore pattern",
		Long:  `Remove ignore pattern.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println(`Usage:
  dotfiles-collector ignore remove <pattern>`)
				return
			}
			err := app.RemoveIgnorePattern(args[0])
			if err != nil {
				fmt.Printf("Failed to remove pattern: %s\n", err)
				return
			}
		},
	}

	listIgnorePatterns := &cobra.Command{
		Use:   "list",
		Short: "List ignore patterns",
		Long:  "List ignore patterns added to the collector.",
		Run: func(cmd *cobra.Command, args []string) {
			var sb strings.Builder
			patterns, err := app.GetIgnorePatterns()
			if err != nil {
				fmt.Printf("Failed to get ignore patterns: %s\n", err)
				return
			}
			for _, pattern := range patterns {
				sb.WriteString(pattern + "\n")
			}
			fmt.Print(sb.String())
		},
	}

	rootCmd.AddCommand(ignoreCmd)
	ignoreCmd.AddCommand(addPattern)
	ignoreCmd.AddCommand(removePattern)
	ignoreCmd.AddCommand(listIgnorePatterns)
}
