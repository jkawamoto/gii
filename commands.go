package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// GlobalFlags defines global flags.
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "GOPATH",
		Name:   "gopath",
		Usage:  "`GOPATH`",
	},
}

// Commands defines sub commands.
var Commands = []cli.Command{}

// CommandNotFound defines an action when a given command won't be supported.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(
		os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
