-- name: GetCollectPaths :many
SELECT * FROM collect_paths;

-- name: AddCollectPath :exec
INSERT INTO collect_paths (path, parent_dir) VALUES (?, ?);

-- name: RemoveCollectPath :exec
DELETE FROM collect_paths WHERE id = ?;

-- name: GetIgnoreRegexps :many
SELECT * FROM ignore_patterns;

-- name: AddIgnorePath :exec
INSERT INTO ignore_patterns (pattern) VALUES (?);

-- name: RemoveIgnoreRegexp :exec
DELETE FROM ignore_patterns WHERE id = ?;
