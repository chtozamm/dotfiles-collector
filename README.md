# Dotfiles Collector

A CLI tool that collects configuration files. It copies files from specified sources to a defined destination (default: `~/dotfiles`).

Currently developed only for Windows, but could be easily extended by tweaking `App` fields in `main.go` and removing the restriction on runtime check.

It uses embedded SQLite3 to store paths to collect and regexp to ignore; database file `dotfile.db` is stored at `App.DATA_DIR`.

## Planned features

- [Cobra](https://github.com/spf13/cobra) CLI commander
- Git execution with [go-git](https://github.com/go-git/go-git)
