CREATE TABLE collect_paths (
  id         INTEGER PRIMARY KEY,
  path       TEXT NOT NULL UNIQUE,
	parent_dir TEXT NOT NULL,
  created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%fZ', 'now'))
);

CREATE TABLE ignore_patterns (
  id         INTEGER PRIMARY KEY,
  pattern    TEXT NOT NULL UNIQUE,
  created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%fZ', 'now'))
);
