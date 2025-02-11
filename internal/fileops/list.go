package fileops

import (
	"cmp"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// File represents a single file or a directory.
type File struct {
	Path     string
	Children []File
	IsDir    bool
}

// ListFiles returns a list of collected files sorted by their paths.
func ListFiles(dir string) ([]File, error) {
	var files []File

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(dir, entry.Name())

		if filepath.Base(entryPath) == ".git" {
			continue
		}

		file := File{Path: entryPath}

		if entry.IsDir() {
			// Recursively list files in subdirectories
			subFiles, err := ListFiles(entryPath)
			if err != nil {
				return nil, err
			}
			file.IsDir = true
			file.Children = subFiles
		}

		files = append(files, file)
	}

	slices.SortFunc(files, func(a, b File) int {
		if a.IsDir && !b.IsDir {
			return -1 // Directories come before non-directories
		}
		if !a.IsDir && b.IsDir {
			return 1 // Non-directories come after directories
		}
		return cmp.Compare(strings.ToLower(a.Path), strings.ToLower(b.Path))
	})

	return files, nil
}
