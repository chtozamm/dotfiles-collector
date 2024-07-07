package main

import (
	"log"

	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

// copyFiles copies files from the source paths to the destination
// specified in the application.
func (app *App) copyFiles() {
	for _, src := range app.SourcePaths {
		if err := fileops.Copy(app.Destination, src, app.BufferSize, app.IgnorePatterns); err != nil {
			log.Printf("Error copying %s: %v", src.Path, err)
		}
	}
}
