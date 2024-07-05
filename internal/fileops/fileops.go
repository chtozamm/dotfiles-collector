// Package fileops provides functions for file and directory operations.
// At the moment, copying is the only implemented operation.
package fileops

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Source represents the source file or directory to be copied,
// including its path and optional parent directory.
type Source struct {
	Path      string
	ParentDir string
}

// Copy copies files and directories from source (src) to destination (dst).
// It skips paths listed in ignorePaths and includes src's parent directory (if provided) in dst.
//
// WARNING: Existing files and directories with the same names will be overwritten.
func Copy(dst string, src Source, bufferSize int, ignorePaths map[string]bool) error {
	sourceFileStat, err := os.Stat(src.Path)
	if err != nil {
		return fmt.Errorf("cannot stat source file %s: %v", src.Path, err)
	}

	if src.ParentDir != "" {
		dst = filepath.Join(dst, src.ParentDir)
	}

	if sourceFileStat.IsDir() {
		return copyDirContents(src.Path, dst, bufferSize, ignorePaths)
	}
	return copyFile(src.Path, dst, bufferSize, ignorePaths)
}

// copyDir copies the directory from source (src) to destination (dst).
// It returns early if the src path is listed in ignorePaths. It creates dst directory with src's base name.
//
// WARNING: Existing files and directories with the same names will be overwritten.
func copyDir(src, dst string, bufferSize int, ignorePaths map[string]bool) error {
	if shouldIgnoreSource(src, ignorePaths) {
		return nil
	}

	currDirName := filepath.Base(src)
	dstDir := filepath.Join(dst, currDirName)

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("cannot stat source directory %s: %v", src, err)
	}

	// Remove directory if it already exists
	if _, err := os.Stat(dstDir); err == nil {
		if err = os.RemoveAll(dstDir); err != nil {
			return fmt.Errorf("error removing existing directory %s: %v", dstDir, err)
		}
	}

	// Create directory preserving permissions
	if err := os.MkdirAll(dstDir, sourceFileStat.Mode().Perm()); err != nil {
		return fmt.Errorf("error creating directory %s: %v", dstDir, err)
	}

	items, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %v", src, err)
	}

	for _, item := range items {
		srcPath := filepath.Join(src, item.Name())
		if item.IsDir() {
			if err := copyDir(srcPath, dstDir, bufferSize, ignorePaths); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstDir, bufferSize, ignorePaths); err != nil {
				return err
			}
		}
	}

	return nil
}

// copyDirContents copies contents of directory from source (src) to destination (dst).
// It returns early if the src is listed in ignorePaths. Unlike copyDir, it does not create a new directory for src.
//
// WARNING: Existing files and directories with the same names will be overwritten.
func copyDirContents(src, dst string, bufferSize int, ignorePaths map[string]bool) error {
	if shouldIgnoreSource(src, ignorePaths) {
		return nil
	}

	items, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %v", src, err)
	}

	for _, item := range items {
		srcPath := filepath.Join(src, item.Name())
		if item.IsDir() {
			if err := copyDir(srcPath, dst, bufferSize, ignorePaths); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dst, bufferSize, ignorePaths); err != nil {
				return err
			}
		}
	}

	return nil
}

// copyFile copies the file from source (src) to destination (dst).
// It ensures dst directory exists before copying. It returns early if the file is listed in ignorePaths.
//
// WARNING: Existing file with the same name in dst will be overwritten.
func copyFile(src, dst string, bufferSize int, ignorePaths map[string]bool) error {
	if shouldIgnoreSource(src, ignorePaths) {
		return nil
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("cannot stat source file %s: %v", src, err)
	}

	dstPath := filepath.Join(dst, filepath.Base(src))

	// Create parent directory if it doesn't exist
	dstDir := filepath.Dir(dstPath)
	if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directory %s: %v", dstDir, err)
	}

	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source file %s: %v", src, err)
	}
	defer source.Close()

	destination, err := os.OpenFile(dstPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, sourceFileStat.Mode().Perm())
	if err != nil {
		return fmt.Errorf("error creating destination file %s: %v", dstPath, err)
	}
	defer destination.Close()

	buf := make([]byte, bufferSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("error reading from source file %s: %v", src, err)
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return fmt.Errorf("error writing to destination file %s: %v", dstPath, err)
		}
	}

	return nil
}

// shouldIgnoreSource checks if the source path (src) should be ignored.
func shouldIgnoreSource(src string, ignorePaths map[string]bool) bool {
	cleanSrc := filepath.Clean(src)
	for ignorePath := range ignorePaths {
		if strings.HasPrefix(cleanSrc, filepath.Clean(ignorePath)) {
			return true
		}
	}
	return false
}
