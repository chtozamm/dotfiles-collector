package app

import (
	"github.com/chtozamm/dotfiles-collector/internal/database"
)

// SourcePath represents a path to collect files from.
type SourcePath struct {
	ID     int64
	Path   string
	Subdir string
}

// Application is the heart of the Dotfiles Collector application.
type Application struct {
	DB          *database.Queries // Interface for executing database queries.
	DataDir     string            // Directory where the application data is stored.
	Destination string            // Directory where the collected files are copied.
	Name        string            // Name of the application.
}

// New returns a new instance of the application with the provided name.
func New(name string) *Application {
	return &Application{Name: name}
}
