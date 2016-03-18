package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Usage = "Command line access to your adjust data."

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound
	app.Action = DefaultAction

	app.Run(os.Args)
}
