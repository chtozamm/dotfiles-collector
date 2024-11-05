package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/chtozamm/dotfiles-collector/internal/database"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

// copyFiles copies files from the source paths to the destination
// specified in the application.
func (app *App) CopyFiles() error {
	if len(app.SourcePaths) == 0 {
		return fmt.Errorf("no source paths found in database to collect")
	}
	for _, src := range app.SourcePaths {
		if err := fileops.Copy(app.Destination, src, app.BufferSize, app.IgnorePatterns); err != nil {
			return fmt.Errorf("failed to copy %s: %v\n", src.Path, err)
		}
	}
	return nil
}

// listCollectedFiles lists the currently collected files in the destination directory.
func (app *App) ListCollectedFiles() {
	files, err := fileops.ListFiles(app.Destination, 1)
	if err != nil {
		log.Printf("Error listing files in destination %s: %v", app.Destination, err)
		return
	}
	if len(files) == 0 {
		fmt.Println("No files collected in destination")
		return
	}
	fmt.Println("Collected files:")
	for _, file := range files {
		fmt.Println(file)
	}
}

// listPaths lists the paths added to the collector.
func (app *App) ListPaths() {
	if len(app.SourcePaths) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintln(w, "INDEX\tPATH\tPARENT_DIR\tID")
		for i, path := range app.SourcePaths {
			fmt.Fprintf(w, "%d\t%s\t%s\t%d\n", i+1, path.Path, path.ParentDir, path.ID)
		}
		w.Flush()
	} else {
		fmt.Println("No source paths to collect added")
	}
}

// listIgnorePatterns lists the regular expressions patterns that will be ignored by the collector.
func (app *App) ListIgnorePatterns() {
	if len(app.SourcePaths) > 0 {
		fmt.Println("Patterns to ignore:")
		i := 1
		for pattern := range app.IgnorePatterns {
			fmt.Printf("%d. %s\n", i, pattern)
			i++
		}
	} else {
		fmt.Println("No ignore patterns added")
	}
}

// addPath adds a new source path to the collector.
func (app *App) AddPath(path, parentDir string) {
	err := app.DB.AddCollectPath(context.Background(), database.AddCollectPathParams{Path: path, ParentDir: parentDir})
	if err != nil {
		log.Printf("Error adding path %s: %v", path, err)
		return
	}
	if parentDir != "" {
		fmt.Println("Path added:", path, "ParentDir:", parentDir)
	} else {
		fmt.Println("Path added:", path)
	}
}

// removePath removes a source path from the collector.
func (app *App) RemovePath(pathID int64) {
	err := app.DB.RemoveCollectPath(context.Background(), pathID)
	if err != nil {
		log.Printf("Error removing path with ID %d: %v", pathID, err)
		return
	}
	fmt.Println("Path removed:", pathID)
}

// addIgnorePattern adds a new source path to the collector.
func (app *App) AddIgnorePattern(pattern string) {
	err := app.DB.AddIgnorePattern(context.Background(), pattern)
	if err != nil {
		log.Printf("Error adding ignore pattern %s: %v", pattern, err)
		return
	}
	fmt.Println("Pattern added:", pattern)
}

// listConfig lists the application settings.
func (app *App) ListConfig() {
	config, _ := app.DB.GetAppConfig(context.Background())
	if config == nil {
		fmt.Println("Configuration was not set")
		return
	}
	for k, v := range config {
		fmt.Printf("%v: %v\n", k, v)
	}
}
