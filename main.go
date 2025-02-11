package main

import (
	"fmt"
	"os"

	"github.com/chtozamm/dotfiles-collector/internal/app"
	"github.com/chtozamm/dotfiles-collector/internal/cli"
	"github.com/chtozamm/dotfiles-collector/internal/tui"

	_ "github.com/mattn/go-sqlite3"
)

func run(f func() error, action string) {
	err := f()
	if err != nil {
		fmt.Printf("Failed to %s: %v\n", action, err)
		os.Exit(1)
	}
}

func main() {
	app := app.New("dotfiles-collector")

	run(app.SetupDataDir, "set up directory for storing data")
	run(app.SetupDB, "set up database")

	if len(os.Args) == 1 {
		// Start terminal application (TUI) if no additional arguments are provided
		if err := tui.Execute(app); err != nil {
			fmt.Printf("Failed to start terminal interface: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Otherwise run command-line interface
		if err := cli.Execute(app); err != nil {
			fmt.Printf("Failed to start command-line application: %v\n", err)
			os.Exit(1)
		}
	}
}
