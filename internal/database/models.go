// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

type CollectPath struct {
	ID        int64
	Path      string
	ParentDir string
	CreatedAt string
}

type IgnoreRegexp struct {
	ID        int64
	Regexp    string
	CreatedAt string
}
