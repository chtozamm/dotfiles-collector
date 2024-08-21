package app

import (
	"github.com/chtozamm/dotfiles-collector/internal/database"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

// App represents the configuration of the Dotfiles Collector application.
type App struct {
	BufferSize     int               // Size of the buffer used for file operations.
	DB             *database.Queries // Interface for executing database queries.
	DataDir        string            // Directory where application data is stored, including the database file.
	Destination    string            // Directory where collected files are copied.
	IgnorePatterns map[string]bool   // Map of regular expressions indicating files to ignore during collection.
	Name           string            // Name of the application.
	SourcePaths    []fileops.Source  // Slice of file paths or sources from which files are collected.
}

func New(name string, bufferSize int) *App {
	return &App{
		Name:       name,
		BufferSize: bufferSize,
	}
}
