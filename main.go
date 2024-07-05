package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/chtozamm/dotfiles-collector/internal/fileops"
)

type App struct {
	BufferSize   int
	UserProfile  string
	AppData      string
	LocalAppData string
	Destination  string
	SourcePaths  []fileops.Source
	IgnorePaths  map[string]bool
}

// InitializePaths initializes system-specific paths.
func (app *App) InitializePaths() {
	app.UserProfile = os.Getenv("USERPROFILE")
	app.AppData = os.Getenv("APPDATA")
	app.LocalAppData = os.Getenv("LOCALAPPDATA")
	app.Destination = filepath.Join(app.UserProfile, "dotfiles")
}

// SetupCollectPaths sets up the source paths to be collected
func (app *App) SetupCollectPaths() {
	app.SourcePaths = []fileops.Source{
		{Path: filepath.Join(app.AppData, "Code", "User"), ParentDir: "vscode"},
		{Path: filepath.Join(app.LocalAppData, `nvim`), ParentDir: "nvim"},
		{Path: filepath.Join(app.UserProfile, `.gitconfig`), ParentDir: ""},
		{Path: filepath.Join(app.UserProfile, "iCloudDrive", "iCloud~md~obsidian", "chtozamm", ".obsidian.vimrc"), ParentDir: "Obsidian"},
	}
}

// SetupCollectPaths sets up the source paths to be collected
func (app *App) SetupIgnorePaths() {
	app.IgnorePaths = map[string]bool{
		filepath.Join(app.AppData, "Code", "User", `globalStorage`):    true,
		filepath.Join(app.AppData, "Code", "User", `History`):          true,
		filepath.Join(app.AppData, "Code", "User", `sync`):             true,
		filepath.Join(app.AppData, "Code", "User", `workspaceStorage`): true,
	}
}

func main() {
	// Only supoprt Windows for now
	if runtime.GOOS != "windows" {
		log.Fatal("Only Windows is currently supported.")
	}

	app := &App{BufferSize: 4096}
	app.InitializePaths()
	app.SetupCollectPaths()
	app.SetupIgnorePaths()

	// Copy files from source to destination
	for _, src := range app.SourcePaths {
		if err := fileops.Copy(app.Destination, src, app.BufferSize, app.IgnorePaths); err != nil {
			log.Printf("Error copying %s: %v", src.Path, err)
		}
	}
}
