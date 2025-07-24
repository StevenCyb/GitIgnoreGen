package main

import (
	"fmt"
	"os"
	"time"

	"github.com/StevenCyb/GitIgnoreGen/pkg/git"
	"github.com/StevenCyb/GitIgnoreGen/pkg/handler"
	"github.com/StevenCyb/GoCLI/pkg/cli"
)

const timeout = time.Duration(5 * time.Second)
const templateURL = "https://github.com/StevenCyb/GitIgnoreGen/tree/main/templates"

func main() {
	client, err := git.New(templateURL)
	if err != nil {
		fmt.Printf("Error creating Git client: %v\n", err)
		os.Exit(1)
	}

	args := os.Args
	var c *cli.CLI
	c = cli.New(
		cli.Name("GitignoreGen"),
		cli.Banner(`╔═╗┬┌┬┐┬┌─┐┌┐┌┌─┐┬─┐┌─┐╔═╗┌─┐┌┐┌
║ ╦│ │ ││ ┬││││ │├┬┘├┤ ║ ╦├┤ │││
╚═╝┴ ┴ ┴└─┘┘└┘└─┘┴└─└─┘╚═╝└─┘┘└┘`),
		cli.Description("A CLI tool to generate .gitignore for various cases."),
		cli.Version("1.0.0"),
		cli.Command("list",
			cli.Description("List available .gitignore templates."),
			cli.Example("cli list"),
			cli.Handler(handler.ListHandler(client, timeout)),
		),
		cli.Command(
			"build",
			cli.Description("Build a .gitignore file in the current working directory."),
			cli.Example("cli build golang macos ..."),
			cli.Handler(handler.BuildHandler(client, timeout, args)),
		),
		cli.Command(
			"update",
			cli.Description("Update a .gitignore file at current working directory."),
			cli.Example("cli update"),
			cli.Handler(handler.UpdateHandler(client, timeout)),
		),
		cli.Command(
			"version",
			cli.Description("Get the version of the CLI."),
			cli.Example("cli version"),
			cli.Handler(
				func(_ *cli.Context) error {
					fmt.Println("1.0.0")
					return nil
				},
			),
		),
		cli.Command(
			"help",
			cli.Description("Show help information."),
			cli.Example("cli help"),
			cli.Handler(func(_ *cli.Context) error {
				c.PrintHelp()
				return nil
			},
			),
		),
	)

	if _, err = c.RunWith(args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
