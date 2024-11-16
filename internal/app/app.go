package app

import (
	"github.com/chtozamm/dotfiles-collector/internal/database"
)

// Source represents a source path entry.
type Source struct {
	ID        int64
	Path      string
	ParentDir string
}

// App represents the configuration of the Dotfiles Collector application.
type App struct {
	DB             *database.Queries // Interface for executing database queries.
	DataDir        string            // Directory where application data is stored, including the database file.
	Destination    string            // Directory where collected files are copied.
	IgnorePatterns map[string]bool   // Map of regular expressions indicating files to ignore during collection.
	Name           string            // Name of the application.
	SourcePaths    []Source          // Slice of file paths or sources from which files are collected.
}

func New() *App {
	return &App{Name: "dotfiles-collector"}
}
