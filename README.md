# Dotfiles Collector

A CLI tool that collects configuration files. It copies files from specified sources to a defined destination (default: `~/dotfiles`). Supports both Windows and Unix-like operating systems.

It uses embedded SQLite3 to store paths to collect and regexp to ignore; database file `dotfile.db` is stored at `App.DATA_DIR`.

## Planned features

- [Cobra](https://github.com/spf13/cobra) CLI commander
- Git execution with [go-git](https://github.com/go-git/go-git)
- If no command-line arguments provided, it starts TUI
- Command-line arguments for listing, adding and removing items from database
