package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	configureCLI()

	app := cli.NewApp()
	app.Name = Name
	app.HelpName = Name
	app.Version = Version
	app.Usage = "Command line access to your adjust data."

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
