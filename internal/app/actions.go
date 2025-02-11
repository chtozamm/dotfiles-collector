package app

import (
	"cmp"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/chtozamm/dotfiles-collector/internal/database"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

// CopyFiles prepares source paths and copies them to the destination.
func (app *Application) CopyFiles() error {
	paths, err := app.GetCollectPaths()
	if err != nil {
		return fmt.Errorf("get paths: %v", err)
	}

	if len(paths) == 0 {
		return fmt.Errorf("no paths found in database")
	}

	ignorePatterns, err := app.GetIgnorePatterns()
	if err != nil {
		return fmt.Errorf("get ignore patterns: %v", err)
	}

	for _, src := range paths {
		dstPath := app.Destination
		// Append parent directory name to destination if specified
		if src.Subdir != "." {
			dstPath = filepath.Join(app.Destination, src.Subdir)
		}
		if err := fileops.Copy(src.Path, dstPath, true, true, ignorePatterns); err != nil {
			return fmt.Errorf("copy %s: %v", src.Path, err)
		}
	}
	return nil
}

// GetCollectPaths returns a list of source paths added to the collector.
func (app *Application) GetCollectPaths() ([]SourcePath, error) {
	paths := []SourcePath{}
	collectPaths, err := app.DB.GetCollectPaths(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get collect paths: %v", err)
	}
	for _, path := range collectPaths {
		paths = append(paths, SourcePath{ID: path.ID, Path: path.Path, Subdir: path.ParentDir})
	}

	slices.SortFunc(paths, func(a, b SourcePath) int {
		return cmp.Compare(a.Path, b.Path)
	})

	return paths, nil
}

// GetIgnorePatterns returns a list of ignore patterns added to the collector.
func (app *Application) GetIgnorePatterns() ([]string, error) {
	patterns := []string{}
	patternsEntries, err := app.DB.GetIgnorePatterns(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get ignore patterns: %v", err)
	}
	for _, pattern := range patternsEntries {
		patterns = append(patterns, pattern.Pattern)
	}

	slices.SortFunc(patterns, func(a, b string) int {
		return cmp.Compare(a, b)
	})

	return patterns, nil
}

// AddCollectPath adds a new path to the collector.
func (app *Application) AddCollectPath(path, parentDir string) error {
	if parentDir == "" {
		// Check for "->" in the given path: if it is provided,
		// the left side is path, the right side is parentDir
		matches := regexp.MustCompile(`\s*([^->]+?)\s*->\s*([^->]+)`).FindStringSubmatch(path)
		if len(matches) == 3 {
			path, parentDir = strings.TrimSpace(matches[1]), strings.TrimSpace(matches[2])
		}

		// Check for environmental variables and replace them with values
		matches = regexp.MustCompile(`(\$\w+)`).FindStringSubmatch(path)
		if len(matches) > 0 {
			for _, match := range matches[1:] {
				env, found := os.LookupEnv(match[1:])
				if !found {
					return fmt.Errorf("path contains env %s, but it cannot be retrieved", match)
				}
				path = strings.ReplaceAll(path, match, env)
			}
		}
	}

	// Clean given paths, trim quotes
	path = filepath.Clean(strings.Trim(path, "'\""))
	parentDir = filepath.Clean(strings.Trim(parentDir, "'\""))
	if parentDir == "." {
		parentDir = ""
	}

	// Check if the path exists
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path does not exist: %s", path)
		}
		return fmt.Errorf("check path %s: %v", path, err)
	}

	// Get the absolute path with correct case
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("get absolute path for %s: %v", path, err)
	}

	// Resolve symlinks to ensure correct case
	resolvedPath, err := filepath.EvalSymlinks(absolutePath)
	if err != nil {
		return fmt.Errorf("resolve symlinks for %s: %v", absolutePath, err)
	}
	path = resolvedPath

	// Check if path already added
	_, err = app.DB.GetCollectPath(context.Background(), path)
	if err == nil {
		return fmt.Errorf("path %s already exists", path)
	}

	// Add the path to the database
	err = app.DB.AddCollectPath(context.Background(), database.AddCollectPathParams{Path: path, ParentDir: parentDir})
	if err != nil {
		return fmt.Errorf("add path %s: %v", path, err)
	}

	return nil
}

// AddIgnorePattern adds an ignore pattern to the collector.
func (app *Application) AddIgnorePattern(pattern string) error {
	// Try to compile regex before proceeding
	_, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	// Check if pattern already added
	_, err = app.DB.GetIgnorePattern(context.Background(), pattern)
	if err == nil {
		return fmt.Errorf("pattern %s already exists", pattern)
	}

	err = app.DB.AddIgnorePattern(context.Background(), pattern)
	if err != nil {
		return fmt.Errorf("add pattern %s: %v", pattern, err)
	}
	return nil
}

// RemoveCollectPath removes a source path from the collector.
func (app *Application) RemoveCollectPath(pathname string) error {
	_, err := app.DB.GetCollectPath(context.Background(), pathname)
	if err != nil {
		return fmt.Errorf("path %s does not exist", pathname)
	}
	err = app.DB.RemoveCollectPath(context.Background(), pathname)
	if err != nil {
		return fmt.Errorf("remove path %s: %v", pathname, err)
	}
	return nil
}

// RemoveIgnorePattern removes an ignore pattern from the collector.
func (app *Application) RemoveIgnorePattern(patternString string) error {
	_, err := app.DB.GetIgnorePattern(context.Background(), patternString)
	if err != nil {
		return fmt.Errorf("pattern %s does not exist", patternString)
	}
	err = app.DB.RemoveIgnorePattern(context.Background(), patternString)
	if err != nil {
		return fmt.Errorf("remove pattern %s: %v", patternString, err)
	}
	return nil
}

// RemoveCollectPaths removes the given paths from the database.
func (app *Application) RemoveCollectPaths(pathnames []string) error {
	for _, pathname := range pathnames {
		err := app.RemoveCollectPath(pathname)
		if err != nil {
			return err
		}
	}
	return nil
}

// RemoveIgnorePatterns removes the given patterns from the database.
func (app *Application) RemoveIgnorePatterns(patterns []string) error {
	for _, pattern := range patterns {
		err := app.RemoveIgnorePattern(pattern)
		if err != nil {
			return err
		}
	}
	return nil
}
