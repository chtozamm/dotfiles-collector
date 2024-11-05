package cmd

import (
	"log"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dotfiles-collector",
	Short: "A CLI tool for managing configuration files",
	Long: `A CLI tool designed to collect configuration files from specified sources and organize them in a defined destination directory. 
It supports regular expression patterns to ignore unwanted files and directories.`,
}

func Execute(app *app.App) {
	// Cobra configuration
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Add commands
	setupListCmd(app, rootCmd)
	setupPathsCmd(app, rootCmd)
	setupIgnoreCmd(app, rootCmd)
	setupCollectCmd(app, rootCmd)
	// setupConfigCmd(app, rootCmd)

	// Execute commands
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
