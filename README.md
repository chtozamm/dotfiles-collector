# Dotfiles Collector

Dotfiles Collector is a command-line interface (CLI) tool designed to collect configuration files from specified sources and organize them in a defined destination directory. It supports both Windows and Unix-like operating systems.

## Features

- **File Collection**: Copies configuration files from sources defined in the database to the `App.Destination` directory (`~/dotfiles` by default).
- **SQLite Database**: Manages paths for file collection and regular expressions for file exclusion using an embedded SQLite3 database (`dotfiles.db`), stored at `App.DataDir`.
- **Planned Features**:
  - Integration with [Cobra](https://github.com/spf13/cobra) CLI framework for command-line operations.
  - Git integration using [go-git](https://github.com/go-git/go-git) for version control of collected files.
  - Interactive mode: Starts a user-friendly text-based user interface (TUI) if no command-line arguments are provided.
  - Additional command-line arguments allow for listing, adding, and removing source paths and ignore patterns from the database.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
