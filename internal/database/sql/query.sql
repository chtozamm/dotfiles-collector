-- name: GetCollectPaths :many
SELECT * FROM collect_paths;

-- name: AddCollectPath :exec
INSERT INTO collect_paths (path, parent_dir) VALUES (?, ?);

-- name: GetIgnorePaths :many
SELECT * FROM ignore_paths;

-- name: AddIgnorePath :exec
INSERT INTO ignore_paths (regexp) VALUES (?);
