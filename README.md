# GitIgnoreGen

This CLI tool generates .gitignore for various cases. This is a fun project to help developers quickly set up their projects with the appropriate ignore files.
Feel free to contribute or suggest improvements!

Also check out the [gitignore.io](https://www.toptal.com/developers/gitignore).

## Installation
```bash
go install github.com/StevenCyb/ServMock@latest
```

## Usage
```bash
╔═╗┬┌┬┐┬┌─┐┌┐┌┌─┐┬─┐┌─┐╔═╗┌─┐┌┐┌
║ ╦│ │ ││ ┬││││ │├┬┘├┤ ║ ╦├┤ │││
╚═╝┴ ┴ ┴└─┘┘└┘└─┘┴└─└─┘╚═╝└─┘┘└┘

1.0.0

A CLI tool to generate .gitignore for various cases.

Usage: 
        GitignoreGen <command>

Commands:
        list
                List available .gitignore templates.
                Example: cli list
        build
                Build a .gitignore file in the current working directory.
                Example: cli build golang macos ...
        update
                Update a .gitignore file at current working directory.
                Example: cli update
        version
                Get the version of the CLI.
                Example: cli version
        help
                Show help information.
                Example: cli help


Use "GitignoreGen <command> --help" for more information about a command.
```
