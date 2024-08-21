-- name: GetCollectPaths :many
SELECT * FROM collect_paths;

-- name: AddCollectPath :exec
INSERT INTO collect_paths (path, parent_dir) VALUES (?, ?);

-- name: RemoveCollectPath :exec
DELETE FROM collect_paths WHERE id = ?;

-- name: GetIgnorePatterns :many
SELECT * FROM ignore_patterns;

-- name: AddIgnorePattern :exec
INSERT INTO ignore_patterns (pattern) VALUES (?);

-- name: RemoveIgnorePattern :exec
DELETE FROM ignore_patterns WHERE id = ?;

-- name: GetAppConfig :many
SELECT * FROM app_config;

