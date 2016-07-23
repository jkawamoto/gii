//
// main.go
//
// Copyright (c) 2016 Junpei Kawamoto
//
// This software is released under the MIT License.
//
// http://opensource.org/licenses/mit-license.php
//

package main

import (
	"os"

	"github.com/jkawamoto/gii/command"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "Junpei Kawamoto"
	app.Email = "kawamoto.junpei@gmail.com"
	app.Usage = "set repositories which doesn't belong golang project to .goimportsignore."

	app.Flags = GlobalFlags
	app.Action = command.CmdRun
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.EnableBashCompletion = true
	app.Copyright = `Copyright (C) 2016  Junpei Kawamoto
This software is released under the MIT License.`

	app.Run(os.Args)
}
