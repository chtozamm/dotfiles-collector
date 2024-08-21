package main

import (
	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/chtozamm/dotfiles-collector/internal/cmd"
)

var bufferSize = 4096

func main() {
	app := app.New("dotfiles-collector", bufferSize)
	app.SetupDirectories()
	app.SetupDB()
	app.SetupCollectPaths()
	app.SetupIgnorePaths()
	// app.SetupConfig()

	cmd.Execute(app)
}
