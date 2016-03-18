package command

import (
	"log"

	"github.com/adjust/adjust-cli/adjust"
	"github.com/codegangsta/cli"
)

func CmdLogout(c *cli.Context) {
	err := adjust.DestroySession(adjust.DefaultConfigFilename)
	if err != nil {
		log.Fatal("You are not logged in.")
	}
}
