package app

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// SetupDataDir sets up directory paths for the application based on the operating system.
func (app *Application) SetupDataDir() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("get user home directory: %v", err)
	}
	app.Destination = filepath.Join(homeDir, "dotfiles")
	if err := os.MkdirAll(app.Destination, 0740); err != nil {
		return fmt.Errorf("create destination directory %q: %v", app.Destination, err)
	}

	// Set application data directory based on the operating system
	if runtime.GOOS == "windows" {
		// On Windows, use %LOCALAPPDATA%/dotfiles-collector
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			return fmt.Errorf("env is not defined: LOCALAPPDATA")
		}
		app.DataDir = filepath.Join(localAppData, app.Name)
	} else {
		// On Unix-like systems (Linux, macOS), use ~/.config/dotfiles
		app.DataDir = filepath.Join(homeDir, ".config", app.Name)
	}

	return nil
}
