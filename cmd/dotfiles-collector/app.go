package main

import (
	"github.com/chtozamm/dotfiles-collector/internal/database"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

type App struct {
	DB          *database.Queries
	NAME        string
	DATA_DIR    string
	BufferSize  int
	Destination string
	SourcePaths []fileops.Source
	IgnorePaths map[string]bool
}
