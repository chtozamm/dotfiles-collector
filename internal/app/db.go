package app

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chtozamm/dotfiles-collector/internal/database"
)

// SetupDB opens the database connection, initializes the schema
// and assigns the connection to the application.
func (app *Application) SetupDB() error {
	// Ensure the application data directory exists, create if it doesn't
	if err := os.MkdirAll(app.DataDir, 0o740); err != nil {
		return fmt.Errorf("create application data directory: %v", err)
	}

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", filepath.Join(app.DataDir, "dotfiles_collector.db"))
	if err != nil {
		return fmt.Errorf("open database: %v", err)
	}

	// Initialize the database schema (tables and triggers)
	if _, err = db.Exec(schema); err != nil {
		return fmt.Errorf("set up database: %v", err)
	}

	app.DB = database.New(db)

	return nil
}

var schema = `
CREATE TABLE IF NOT EXISTS collect_paths (
  id         INTEGER PRIMARY KEY,
  path       TEXT NOT NULL UNIQUE,
  parent_dir TEXT NOT NULL,
  created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%fZ', 'now'))
);

CREATE TABLE IF NOT EXISTS ignore_patterns (
  id         INTEGER PRIMARY KEY,
  pattern    TEXT NOT NULL UNIQUE,
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
