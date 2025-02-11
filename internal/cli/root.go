package cli

import (
	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dotfiles-collector",
	Short: "An application for managing configuration files",
	Long: `Dotfiles Collector is a command-line interface (CLI) tool designed to collect configuration files 
from specified sources and organize them in a defined destination directory.`,
}

// Execute starts the application in the command mode.
func Execute(app *app.Application) error {
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
		return err
	}

	return nil
}
