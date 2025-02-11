package fileops

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Copy copies a file or a directory to a specified destination.
//
// Optionally, it can overwrite existing files and create destination directory
// if it doesn't exist. If ignore patterns are provided, it can check source against them
// and skip copying if match is found.
func Copy(src, dst string, overwrite bool, createDst bool, ignorePatterns []string) error {
	// Return if source is in the ignore list
	if shouldIgnorePath(src, ignorePatterns) {
		return nil
	}

	// Check if source file exists
	srcFileInfo, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("source %q does not exist", src)
		}
		return fmt.Errorf("stat file: %v", err)
	}

	// Check if destination directory exists
	if !doesDirExist(dst) && !createDst {
		return fmt.Errorf("destination %q does not exist and createDst is set to false", src)
	}

	// Call appropriate copy function
	if srcFileInfo.IsDir() {
		return copyDirectory(src, dst, overwrite, ignorePatterns)
	}
	return copyFile(src, dst, overwrite, ignorePatterns)
}

// copyFile copies a file to a specified destination.
func copyFile(src, dst string, overwrite bool, ignorePatterns []string) error {
	// Return if source is in the ignore list
	if shouldIgnorePath(src, ignorePatterns) {
		return nil
	}

	// Append file name to destination path
	dst = filepath.Join(dst, filepath.Base(src))

	// Check if the destination file exists
	if !overwrite {
		if _, err := os.Stat(dst); err == nil {
			if os.IsExist(err) {
				return fmt.Errorf("destination file %q already exists and overwrite is set to false", dst)
			}
			return fmt.Errorf("stat file: %v", err)
		}
	}

	// Create parent directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(dst), 0740); err != nil {
		return fmt.Errorf("create directory %q: %v", filepath.Dir(dst), err)
	}

	// Open source file
	srcFile, err := os.Open(filepath.Clean(src))
	if err != nil {
		return fmt.Errorf("open source file %q: %v", src, err)
	}
	defer srcFile.Close()

	srcFileInfo, err := srcFile.Stat()
	if err != nil {
		return fmt.Errorf("stat source file %q: %v", src, err)
	}

	// Open or create destination file
	dstFile, err := os.OpenFile(filepath.Clean(dst), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, srcFileInfo.Mode())
	if err != nil {
		return fmt.Errorf("open destination file %q: %v", dst, err)
	}
	defer dstFile.Close()

	// Copy contents from source to destination
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return fmt.Errorf("copy %q to %q: %v", src, dst, err)
	}

	return nil
}

// copyDirectory copies a directory and its contents to a specified destination.
func copyDirectory(src, dst string, overwrite bool, ignorePatterns []string) error {
	// Return if source is in the ignore list
	if shouldIgnorePath(src, ignorePatterns) {
		return nil
	}

	// Append directory name to destination path
	dst = filepath.Join(dst, filepath.Base(src))

	// Check if the source directory exists
	if !doesDirExist(src) {
		return fmt.Errorf("source directory %q does not exist", src)
	}

	// Create the destination directory
	if err := os.MkdirAll(dst, 0740); err != nil {
		return fmt.Errorf("create destination directory %q: %v", dst, err)
	}

	// Read the contents of the source directory
	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("read source directory %q: %v", src, err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		if entry.IsDir() {
			if err := copyDirectory(srcPath, dst, overwrite, ignorePatterns); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dst, overwrite, ignorePatterns); err != nil {
				return err
			}
		}
	}

	return nil
}
