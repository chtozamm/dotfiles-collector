package app

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// SetupDirectories sets up directory paths for the application based on the operating system.
func (app *App) SetupDirectories() {
	// Retrieve the user's home directory.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user's home directory: %v", err)
	}
	app.Destination = filepath.Join(homeDir, "dotfiles")

	// Set application data directory based on the operating system
	if runtime.GOOS == "windows" {
		// On Windows, use %LOCALAPPDATA%/dotfiles-collector
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			log.Fatal("LOCALAPPDATA environment variable is not defined")
		}
		app.DataDir = filepath.Join(localAppData, app.Name)
	} else {
		// On Unix-like systems (Linux, macOS), use ~/.config/dotfiles
		app.DataDir = filepath.Join(homeDir, ".config", app.Name)
	}
}

// SetupCollectPaths sets up the source paths to be collected.
func (app *App) SetupCollectPaths() {
	app.SourcePaths = []Source{}
	collectPaths, _ := app.DB.GetCollectPaths(context.Background())
	for _, path := range collectPaths {
		app.SourcePaths = append(app.SourcePaths, Source{ID: path.ID, Path: path.Path, ParentDir: path.ParentDir})
	}
}

// SetupCollectPaths sets up the source paths to be collected.
func (app *App) SetupIgnorePaths() {
	app.IgnorePatterns = make(map[string]bool)
	ignorePaths, _ := app.DB.GetIgnorePatterns(context.Background())
	for _, path := range ignorePaths {
		app.IgnorePatterns[path.Pattern] = true
	}
}
