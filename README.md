# Dotfiles Collector

Dotfiles Collector is a command-line interface (CLI) tool designed to gather configuration files across an operating system.
It allows you to manage paths to files and directories you want to collect and then copy them in one place with a single command.
Additionally, you can add ignore patterns as regular expressions that Dotfiles Collector would ignore if encountered. 

## Requirements

- Go is required to build or [install](#installation) the application
- Both **Windows** and **Unix**-like operating systems are supported

## Foreword

Dotfiles Collector creates two directories when you use it for the first time:

|             | Collected Files          | Application Data                    |
| ----------- | ------------------------ | ----------------------------------- |
| **Windows** | `%USERPROFILE%/dotfiles` | `%LOCALAPPDATA%/dotfiles-collector` |
| **Unix**    | `~/dotfiles`             | `~/.config/dotfiles-collector`      |

The first directory is where you find the files you have collected. 
The second directory contains an SQLite3 database file with your configuration.

## Usage

> [!NOTE]
> Dotfiles Collector can operate in two modes: **command mode** and **interactive mode**.

### Command Mode

To use **command mode**, run `dotfiles-collector --help` to display available commands.

```plaintext
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

You can specify which directories or files it should collect. 

```sh
dotfiles-collector paths add "~/.gitconfig"
```

It also works with relative to current working directory paths:

```sh
dotfiles-collector paths add .
```

Sometimes you might want the collector to create a subdirectory. To do so, you can specify a second argument 
after the source path. The following example will create a `backup` directory for `sqlite.db`:


```sh
dotfiles-collector paths add "~/my-project/sqlite.db" "backup"
```

Optionally, you can add regular expressions (ignore patterns) that collector will skip if it encounters file 
or directory which name matches the pattern. For example:

```sh
dotfiles-collector ignore add ".*node_modules$"
```

The following Windows example will result in the whole `PowerShell` directory being collected except for its child directory `Modules`:

```sh
dotfiles-collector paths add "D:\Documents\PowerShell"
dotfiles-collector ignore add ".*PowerShell[/\\]Modules$"
dotfiles-collector collect
```

### Interactive mode

To start **interactive mode**, simply run `dotfiles-collector` without additional arguments.

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://i.imgur.com/oSVVtbY.gif">
  <source media="(prefers-color-scheme: light)" srcset="https://i.imgur.com/oSVVtbY.gif">
  <img width="600" alt="A demonstration of interactive mode" src="https://i.imgur.com/oSVVtbY.gif">
</picture>

## Installation

Install with Go:

```sh
go install github.com/chtozamm/dotfiles-collector@latest
```

## Technologies Used

Dotfiles Collector is built using the following tools and libraries:

- **Go**: The programming language used to develop the application.
- **SQLite3**: A lightweight embedded database used to store configuration data.
- **Cobra**: A library for creating command-line applications.
- **Charm**: A library for building interactive terminal applications.

## License

[MIT](https://github.com/chtozamm/dotfiles-collector/blob/main/LICENSE.md)