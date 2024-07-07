-- name: GetCollectPaths :many
SELECT * FROM collect_paths;

-- name: AddCollectPath :exec
INSERT INTO collect_paths (path, parent_dir) VALUES (?, ?);

-- name: RemoveCollectPath :exec
DELETE FROM collect_paths WHERE id = ?;

-- name: GetIgnoreRegexps :many
SELECT * FROM ignore_regexps;

-- name: AddIgnorePath :exec
INSERT INTO ignore_regexps (regexp) VALUES (?);

-- name: RemoveIgnoreRegexp :exec
DELETE FROM ignore_regexps WHERE id = ?;
