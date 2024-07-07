package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

// SetupDirectories sets up directory paths for the application based on the operating system.
func (app *App) setupDirectories() {
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

// SetupCollectPaths sets up the source paths to be collected
func (app *App) setupCollectPaths() {
	app.SourcePaths = []fileops.Source{}
	collectPaths, _ := app.DB.GetCollectPaths(context.Background())
	for _, path := range collectPaths {
		app.SourcePaths = append(app.SourcePaths, fileops.Source{Path: path.Path, ParentDir: path.ParentDir})
	}
	if len(app.SourcePaths) == 0 {
		fmt.Println("No sources found in database to copy")
		return
	}
}

// SetupCollectPaths sets up the source paths to be collected
func (app *App) setupIgnorePaths() {
	app.IgnorePatterns = make(map[string]bool)
	ignorePaths, _ := app.DB.GetIgnoreRegexps(context.Background())
	for _, path := range ignorePaths {
		app.IgnorePatterns[path.Pattern] = true
	}
}
