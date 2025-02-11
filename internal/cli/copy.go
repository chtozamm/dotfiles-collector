package cli

import (
	"fmt"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

func setupCollectCmd(app *app.Application, rootCmd *cobra.Command) {
	collectCmd := &cobra.Command{
		Use:   "collect",
		Short: "Collect files specified in source paths",
		Long:  "Collect files specified in source paths.",
		Run: func(cmd *cobra.Command, args []string) {
			err := app.CopyFiles()
			if err != nil {
				fmt.Printf("Failed to collect files: %v\n", err)
				return
			}
			fmt.Println("Successfully collected the files.")
		},
	}

	rootCmd.AddCommand(collectCmd)
}
