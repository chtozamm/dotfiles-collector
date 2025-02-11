# Dotfiles Collector

Dotfiles Collector is a command-line interface (CLI) tool with a terminal user interface (TUI) designed to gather configuration files from across your operating system into one centralized location. It simplifies the process of backing up your files by allowing you to manage the paths from which you want to collect files and then copy them with a single command.

Additionally, you can specify ignore patterns using regular expressions, which Dotfiles Collector will respect by ignoring any files that match these patterns.

## Requirements

- Go is required to build or [install](#installation) the application
- Both **Windows** and **Unix**-like operating systems are supported

## Directories

Dotfiles Collector creates two directories when you use it for the first time:

|             | Collected Files         | Application Data                   |
| ----------- | ----------------------- | ---------------------------------- |
| **Windows** | `$USERPROFILE/dotfiles` | `$LOCALAPPDATA/dotfiles_collector` |
| **Unix**    | `$HOME/dotfiles`        | `$HOME/.config/dotfiles_collector` |

The first directory is where you will find the files you have collected. The second directory contains an SQLite3 database file that stores your application data.

## Usage

> [!NOTE]
> Dotfiles Collector can operate in two modes: **command mode** and **interactive mode**.

### Interactive mode

To start **interactive mode**, run `dotfiles-collector` without additional arguments:

```sh
dotfiles-collector
```

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="/docs/tui-demo.gif">
  <source media="(prefers-color-scheme: light)" srcset="/docs/tui-demo.gif">
  <img width="600" alt="A demonstration of interactive mode" src="/docs/tui-demo.gif">
</picture>

### Command Mode

To use **command mode**, run `dotfiles-collector --help` to display available commands:

```plaintext
> dotfiles-collector --help

Usage:
  dotfiles-collector [command]

Available Commands:
  collect     Collect files specified in source paths
  list        List collected files
  help        Help about any command
  paths       Manage source paths
  ignore      Manage ignore patterns
```

#### Examples

You can specify which directories or files Dotfiles Collector should collect using the following command:

```sh
dotfiles-collector paths add "$HOME/.gitconfig"
```

It also works with paths relative to the current working directory:

```sh
dotfiles-collector paths add .
```

If you want the collector to create a subdirectory, you can specify a second argument after the source path. The following example will create a `backup` directory for `sqlite.db`:

```sh
dotfiles-collector paths add "$HOME/my-project/sqlite.db" "backup"
```

Optionally, you can add regular expressions (ignore patterns) that the collector will skip if it encounters a file or directory whose name matches the pattern. For example:

```sh
dotfiles-collector ignore add "\.git$"
dotfiles-collector ignore add "node_modules"
```

In the following Windows example, the entire `PowerShell` directory will be collected, except for its child directory `Modules`:

```sh
dotfiles-collector paths add "D:\Documents\PowerShell"
dotfiles-collector ignore add "PowerShell[\/\\]Modules"
dotfiles-collector collect
```

## Installation

You can install Dotfiles Collector using Go:

```sh
go install github.com/chtozamm/dotfiles-collector@latest
```

## Technologies Used

Dotfiles Collector is built using the following tools and libraries:

- **Go**: The programming language used to develop the application.
- **SQLite3**: A lightweight embedded database for storing application data.
- **Cobra**: A library for creating command-line applications.
- **Bubble Tea**: A framework for building interactive terminal applications.

## License

[MIT](/LICENSE.md)
