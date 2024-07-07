package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/chtozamm/dotfiles-collector/internal/database"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS collect_paths (
  id   INTEGER PRIMARY KEY,
  path TEXT NOT NULL UNIQUE,
	parent_dir TEXT NOT NULL,
  created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%fZ', 'now'))
);

CREATE TABLE IF NOT EXISTS ignore_patterns (
  id   INTEGER PRIMARY KEY,
  pattern TEXT NOT NULL UNIQUE,
  created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%fZ', 'now'))
);

CREATE TRIGGER IF NOT EXISTS set_default_parent_dir
BEFORE INSERT ON collect_paths
FOR EACH ROW
WHEN NEW.parent_dir IS NULL
BEGIN
	UPDATE collect_paths
    SET parent_dir = "."
		WHERE id = NEW.id;
END;
`

func (app *App) setupDB() {
	// Ensure the application data directory exists, create it if it doesn't
	if err := os.MkdirAll(app.DataDir, 0755); err != nil {
		log.Fatalf("Error creating application data directory %s: %v", app.DataDir, err)
	}

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", filepath.Join(app.DataDir, "dotfiles.db"))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Initialize the database schema (tables and triggers)
	if _, err = db.Exec(schema); err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	// Assign the database connection to the application
	app.DB = database.New(db)
}
