package main

import (
	"os"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/chtozamm/dotfiles-collector/internal/cmd"
	"github.com/chtozamm/dotfiles-collector/internal/tui"
)

var bufferSize = 4096

func main() {
	app := app.New("dotfiles-collector", bufferSize)
	app.SetupDirectories()
	app.SetupDB()
	app.SetupCollectPaths()
	app.SetupIgnorePaths()
	// app.SetupConfig()

	if len(os.Args) == 1 {
		// Run TUI if no additional arguments are provided
		tui.Execute(app)
	} else {
		// Otherwise behave like CLI
		cmd.Execute(app)
	}
}
