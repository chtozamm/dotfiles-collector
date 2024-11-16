package fileops

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// File represents a single file or a directory.
type File struct {
	Path     string
	Children []File
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
				return nil, err // Return the error directly
			}
			file.Children = subFiles
		}

		files = append(files, file)
	}

	// Case-insensitive sort
	sort.Slice(files, func(i, j int) bool {
		return strings.ToLower(files[i].Path) < strings.ToLower(files[j].Path)
	})

	return files, nil
}
