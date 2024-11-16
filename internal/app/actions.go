package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/chtozamm/dotfiles-collector/internal/database"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

// CopyFiles prepares source paths and copies them to the destination.
func (app *App) CopyFiles() error {
	if len(app.SourcePaths) == 0 {
		return fmt.Errorf("no source paths found in database to collect")
	}
	for _, src := range app.SourcePaths {
		dstPath := app.Destination
		// Append parent directory name to destination if specified
		if src.ParentDir != "." {
			dstPath = filepath.Join(app.Destination, src.ParentDir)
		}
		if err := fileops.Copy(src.Path, dstPath, true, true, app.IgnorePatterns); err != nil {
			return fmt.Errorf("copy %s: %v", src.Path, err)
		}
	}
	return nil
}

// GetCollectPaths returns a list of source paths added to the collector.
func (app *App) GetCollectPaths() []Source {
	paths := make([]Source, len(app.SourcePaths))
	copy(paths, app.SourcePaths)

	// Case-insensitive sort
	sort.Slice(paths, func(i, j int) bool {
		return strings.ToLower(paths[i].Path) < strings.ToLower(paths[j].Path)
	})

	return paths
}

// GetIgnorePatterns returns a list of ignore patterns added to the collector.
func (app *App) GetIgnorePatterns() []string {
	patterns := make([]string, len(app.IgnorePatterns))
	i := 0
	for pattern := range app.IgnorePatterns {
		patterns[i] = pattern
		i++
	}

	// Case-insensitive sort
	sort.Slice(patterns, func(i, j int) bool {
		return strings.ToLower(patterns[i]) < strings.ToLower(patterns[j])
	})

	return patterns
}

// AddCollectPath adds a new path to the collector.
func (app *App) AddCollectPath(path, parentDir string) error {
	// Check if the path exists
	if _, err := os.Stat(path); err == nil {
		// Path exists, get the absolute path with correct case
		absolutePath, err := filepath.Abs(path)
		if err != nil {
			return fmt.Errorf("get absolute path for %s: %v", path, err)
		}

		// Resolve symlinks to ensure correct case
		resolvedPath, err := filepath.EvalSymlinks(absolutePath)
		if err != nil {
			return fmt.Errorf("resolve symlinks for %s: %v", absolutePath, err)
		}

		// Update the path variable to the resolved path
		path = resolvedPath
	} else if os.IsNotExist(err) {
		// Path does not exist, handle accordingly (e.g., return an error or log)
		return fmt.Errorf("path does not exist: %s", path)
	} else {
		// Some other error occurred
		return fmt.Errorf("check path %s: %v", path, err)
	}

	// Check if path already added
	_, err := app.DB.GetCollectPath(context.Background(), path)
	if err == nil {
		return fmt.Errorf("path %s already exists", path)
	}

	// Add the path to the database
	err = app.DB.AddCollectPath(context.Background(), database.AddCollectPathParams{Path: path, ParentDir: parentDir})
	if err != nil {
		return fmt.Errorf("add path %s: %v", path, err)
	}
	if parentDir != "" {
		fmt.Println("Path added:", path, "ParentDir:", parentDir)
	} else {
		fmt.Println("Path added:", path)
	}
	return nil
}

// RemoveCollectPath removes a source path from the collector.
func (app *App) RemoveCollectPath(pathString string) error {
	path, err := app.DB.GetCollectPath(context.Background(), pathString)
	if err != nil {
		return fmt.Errorf("path %s does not exist", pathString)
	}
	err = app.DB.RemoveCollectPath(context.Background(), pathString)
	if err != nil {
		return fmt.Errorf("remove path %s: %v", pathString, err)
	}
	fmt.Printf("Path removed: %s\n", path.Path)
	return nil
}

// AddIgnorePattern adds an ignore pattern to the collector.
func (app *App) AddIgnorePattern(pattern string) error {
	// Check if pattern already added
	_, err := app.DB.GetIgnorePattern(context.Background(), pattern)
	if err != nil {
		return fmt.Errorf("pattern %s already exists", pattern)
	}

	err = app.DB.AddIgnorePattern(context.Background(), pattern)
	if err != nil {
		return fmt.Errorf("add pattern %s: %v", pattern, err)
	}

	fmt.Printf("Pattern added: %s\n", pattern)
	return nil
}

// RemoveIgnorePattern removes an ignore pattern from the collector.
func (app *App) RemoveIgnorePattern(patternString string) error {
	pattern, err := app.DB.GetIgnorePattern(context.Background(), patternString)
	if err != nil {
		return fmt.Errorf("pattern %s does not exist", patternString)
	}
	err = app.DB.RemoveIgnorePattern(context.Background(), patternString)
	if err != nil {
		return fmt.Errorf("remove pattern %s: %v", patternString, err)
	}
	fmt.Printf("Pattern removed: %s\n", pattern.Pattern)
	return nil
}
