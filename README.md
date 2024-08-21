# Dotfiles Collector

Dotfiles Collector is a command-line interface (CLI) tool designed to collect configuration files from specified sources and organize them in a defined destination directory.

## Requirements

Dotfiles Collector supports both Windows and Unix-like operating systems.

- Tested to work with **[Go](https://go.dev/doc/install) v1.22.4** or greater.

## Installation

```shell
go install https://github.com/chtozamm/dotfiles-collector
```

## Usage

To get started, run the command without any command-line arguments to see a list of available actions and options:

```shell
dotfiles-collector
```

This will display a help message detailing the commands you can use. These commands include:

- `dotfiles-collector collect`: Collects configuration files from the defined sources.
- `dotfiles-collector list`: Lists all collected dotfiles.
- `dotfiles-collector paths [add|remove|list]`: Manages source paths for file collection.
- `dotfiles-collector ignore [add|remove|list]`: Manages ignore patterns for file exclusion.
<!--- `dotfiles-collector config`: Configures settings such as the destination path.-->

## Features

- [x] **File Collection**: Copies configuration files from sources defined in the database to the `App.Destination` directory (default: `~/dotfiles`).
- [x] **SQLite Database**: Manages paths for file collection and regular expressions for file exclusion using an embedded SQLite3 database (`dotfiles.db`), stored in `App.DataDir`.
- [x] **Cobra Integration**: Utilizes [Cobra](https://github.com/spf13/cobra) CLI framework for command-line operations.
- [x] **Command-Line Arguments**: Supports additional command-line arguments for listing, adding, and removing source paths and ignore patterns from the database.

**Planned Features**:
  - **Configuration**: Store configuration settings, such as the destination path, in a database table that can be edited using the `config` command.
  <!--- **Interactive mode**: Launch a user-friendly text-based user interface (TUI) when no command-line arguments are provided..-->
