package main

import (
	"github.com/chtozamm/dotfiles-collector/internal/database"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

// App represents the configuration of the Dotfiles Collector application.
type App struct {
	Name           string            // Name of the application.
	DB             *database.Queries // Interface for executing database queries.
	DataDir        string            // Directory where application data is stored, including the database file.
	BufferSize     int               // Size of the buffer used for file operations.
	Destination    string            // Directory where collected files are copied.
	SourcePaths    []fileops.Source  // Slice of file paths or sources from which files are collected.
	IgnorePatterns map[string]bool   // Map of regular expressions indicating files to ignore during collection.
}
