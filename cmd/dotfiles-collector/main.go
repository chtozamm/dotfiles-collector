package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/chtozamm/dotfiles-collector/internal/database"
	"github.com/chtozamm/dotfiles-collector/internal/fileops"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB         *database.Queries
	DATA_DIR   string
	BufferSize int
	// START the following fields should be removed in the future
	UserProfile  string
	AppData      string
	LocalAppData string
	// END
	Destination string
	SourcePaths []fileops.Source
	IgnorePaths map[string]bool
}

func main() {
	// Only supoprt Windows for now
	if runtime.GOOS != "windows" {
		log.Fatal("Only Windows is currently supported.")
	}

	// TODO: change the way localappdata is obtained
	app := &App{BufferSize: 4096, DATA_DIR: filepath.Join(os.Getenv("LOCALAPPDATA"), "dotfiles-collector")}
	app.setupDB()
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

// InitializePaths initializes system-specific paths.
func (app *App) InitializePaths() {
	app.UserProfile = os.Getenv("USERPROFILE")
	app.AppData = os.Getenv("APPDATA")
	app.LocalAppData = os.Getenv("LOCALAPPDATA")
	app.Destination = filepath.Join(app.UserProfile, "dotfiles")
}

// SetupCollectPaths sets up the source paths to be collected
func (app *App) SetupCollectPaths() {
	err := app.DB.AddCollectPath(context.Background(), database.AddCollectPathParams{
		Path:      filepath.Join(app.AppData, "Code", "User"),
		ParentDir: sql.NullString{String: "vscode", Valid: true},
	})
	if err != nil {
		log.Print(err)
	}
	app.DB.AddCollectPath(context.Background(), database.AddCollectPathParams{
		Path: filepath.Join(app.UserProfile, `.gitconfig`),
	})
	app.DB.AddCollectPath(context.Background(), database.AddCollectPathParams{
		Path:      filepath.Join(app.LocalAppData, `nvim`),
		ParentDir: sql.NullString{String: "nvim", Valid: true},
	})
	app.DB.AddCollectPath(context.Background(), database.AddCollectPathParams{
		Path:      filepath.Join(app.UserProfile, "iCloudDrive", "iCloud~md~obsidian", "chtozamm", ".obsidian.vimrc"),
		ParentDir: sql.NullString{String: "Obsidian", Valid: true},
	})

	collectPaths, _ := app.DB.GetCollectPaths(context.Background())

	app.SourcePaths = []fileops.Source{}

	for _, path := range collectPaths {
		// TODO: figure out the way to use path.ParentDir directly without string field
		app.SourcePaths = append(app.SourcePaths, fileops.Source{Path: path.Path, ParentDir: path.ParentDir.String})
	}
}

// SetupCollectPaths sets up the source paths to be collected
func (app *App) SetupIgnorePaths() {
	err := app.DB.AddIgnorePath(context.Background(), `.*Code[/\\]User[/\\]globalStorage$`)
	if err != nil {
		log.Print(err)
	}
	app.DB.AddIgnorePath(context.Background(), `.*Code[/\\]User[/\\]History$`)
	app.DB.AddIgnorePath(context.Background(), `.*Code[/\\]User[/\\]sync$`)
	app.DB.AddIgnorePath(context.Background(), `.*Code[/\\]User[/\\]workspaceStorage$`)

	ignorePaths, _ := app.DB.GetIgnorePaths(context.Background())

	app.IgnorePaths = make(map[string]bool)

	for _, path := range ignorePaths {
		app.IgnorePaths[path.Regexp] = true
	}
}

func (app *App) setupDB() {
	// Create db file if it doesn't exist
	// TODO: check permissions of created directory(-ies)
	if err := os.MkdirAll(app.DATA_DIR, os.ModeDir); err != nil {
		log.Fatalf("Error creating application data directory %s: %v", app.DATA_DIR, err)
	}

	// Connect to the database
	db, err := sql.Open("sqlite3", filepath.Join(app.DATA_DIR, "dotfiles.db"))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Initialize tables
	if _, err = db.Exec(`
CREATE TABLE IF NOT EXISTS collect_paths (
  id   INTEGER PRIMARY KEY,
  path TEXT NOT NULL UNIQUE,
	parent_dir TEXT
);

CREATE TABLE IF NOT EXISTS ignore_paths (
  id   INTEGER PRIMARY KEY,
  regexp TEXT NOT NULL UNIQUE
);
`); err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	app.DB = database.New(db)
}
