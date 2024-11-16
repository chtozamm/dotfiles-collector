package fileops

import (
	"os"
	"path/filepath"
	"regexp"
)

// doesDirExist checks if a directory exists at the given path.
func doesDirExist(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
}

// shouldIgnorePath returns true if the path matches any of the ignore patterns.
func shouldIgnorePath(path string, ignorePatterns map[string]bool) bool {
	cleanPath := filepath.Clean(path)

	for ignorePath := range ignorePatterns {
		re, err := regexp.Compile(ignorePath)
		if err != nil {
			continue // Skip this ignorePath if there's an error
		}
		if re.MatchString(cleanPath) {
			return true
		}
	}

	return false
}
