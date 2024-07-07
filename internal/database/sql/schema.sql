CREATE TABLE collect_paths (
  id   INTEGER PRIMARY KEY,
  path TEXT NOT NULL UNIQUE,
	parent_dir TEXT
);

CREATE TABLE ignore_paths (
  id   INTEGER PRIMARY KEY,
  regexp TEXT NOT NULL UNIQUE
);