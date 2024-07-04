package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

const (
	BUFFERSIZE   = 2048
	DEST_DIRNAME = "dotfiles"
)

var (
	userprofile string
	appdata     string
	localdata   string
	vscode      string
)

func main() {
	// Only supoprt Windows for now
	if runtime.GOOS != "windows" {
		log.Fatal("Only Windows is currently supported.")
	}

	// Initialize paths
	userprofile = os.Getenv("USERPROFILE")
	appdata = os.Getenv("APPDATA")
	localdata = os.Getenv("LOCALAPPDATA")

	destination := path.Join(userprofile, DEST_DIRNAME)
	vscode = filepath.Join(appdata, "Code", "User")

	// Paths to the files and directories to collect
	sourcePaths := []string{
		vscode,
		filepath.Join(localdata, "nvim"),
		filepath.Join(userprofile, ".gitconfig"),
	}

	/* Example of sources to collect:
	nvim
	vscode
	.gitconfig
	.prettier.config.mjs
	.tailwind.config.mjs
	.postcss.config.mjs
	*/

	for _, src := range sourcePaths {
		if err := copy(src, destination); err != nil {
			log.Printf("Error copying %s: %v", src, err)
		}
	}

	// TODO: git-go: commit and push changes to the remote git repository.
	// This feature must be optional. Repository must be initialized in destination and remote must be set.
}

func copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("cannot stat source file %s: %v", src, err)
	}

	if sourceFileStat.IsDir() {
		return copyDir(src, dst)
	}
	return copyFile(src, dst)
}

func copyDir(src, dst string) error {
	if shouldIgnoreSource(src) {
		return nil
	}

	currDirName := filepath.Base(src)
	dstDir := path.Join(dst, currDirName)

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("cannot stat source directory %s: %v", src, err)
	}

	// Remove directory if it already exists
	if _, err := os.Stat(dstDir); err == nil {
		err = os.RemoveAll(dstDir)
		if err != nil {
			return fmt.Errorf("error removing existing directory %s: %v", dstDir, err)
		}
	}

	// Create directory preserving the permissions
	if err := os.Mkdir(dstDir, sourceFileStat.Mode().Perm()); err != nil {
		return fmt.Errorf("error creating directory %s: %v", dstDir, err)
	}

	items, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %v", src, err)
	}

	for _, item := range items {
		srcPath := path.Join(src, item.Name())
		if item.IsDir() {
			if err := copyDir(srcPath, dstDir); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstDir); err != nil {
				return err
			}
		}
	}

	return nil
}

// copyFile copies a file from the source to the destination, preserving permissions.
// It rewrites the file
func copyFile(src, dst string) error {
	if shouldIgnoreSource(src) {
		return nil
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("cannot stat source file %s: %v", src, err)
	}

	dstPath := filepath.Join(dst, filepath.Base(src))

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

	buf := make([]byte, BUFFERSIZE)
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

// shouldIgnoreSource checks if a given source path should be ignored based on predefined rules.
// It currently ignores specific paths related to VS Code settings.
func shouldIgnoreSource(src string) bool {
	var ignorePaths = map[string]bool{
		filepath.Join(vscode, `globalStorage`):    true,
		filepath.Join(vscode, `History`):          true,
		filepath.Join(vscode, `sync`):             true,
		filepath.Join(vscode, `workspaceStorage`): true,
	}
	return ignorePaths[src]
}
