-- name: GetCollectPaths :many
SELECT * FROM collect_paths;

-- name: GetCollectPath :one
SELECT * FROM collect_paths WHERE path = ?;

-- name: AddCollectPath :exec
INSERT INTO collect_paths (path, parent_dir) VALUES (?, ?);

-- name: RemoveCollectPath :exec
DELETE FROM collect_paths WHERE path = ?;

-- name: GetIgnorePatterns :many
SELECT * FROM ignore_patterns;

-- name: GetIgnorePattern :one
SELECT * FROM ignore_patterns WHERE pattern = ?;

-- name: AddIgnorePattern :exec
INSERT INTO ignore_patterns (pattern) VALUES (?);

-- name: RemoveIgnorePattern :exec
DELETE FROM ignore_patterns WHERE pattern = ?;
