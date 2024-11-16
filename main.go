package main

import (
	"os"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/chtozamm/dotfiles-collector/internal/cmd"
	"github.com/chtozamm/dotfiles-collector/internal/tui"
)

func main() {
	app := app.New()
	app.SetupDirectories()
	app.SetupDB()
	app.SetupCollectPaths()
	app.SetupIgnorePaths()

	if len(os.Args) == 1 {
		// Start terminal application (TUI) if no additional arguments are provided
		tui.Execute(app)
	} else {
		// Otherwise run command-line interface
		cmd.Execute(app)
	}
}
