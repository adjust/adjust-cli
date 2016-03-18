package command

import (
	"fmt"
	"log"

	"github.com/adjust/adjust-cli/adjust"
	"github.com/codegangsta/cli"
)

func CmdLogin(c *cli.Context) {
	if c.NArg() > 1 {
		adjust.Fail("SHIT We need only one argument - the user token!")
	}

	if c.NArg() < 1 {
		adjust.Fail(fmt.Sprintf("You need to provide a user token. See '%s login --help'", c.App.Name))
	}

	userToken := c.Args().First()
	session, err := adjust.NewSession(adjust.DefaultConfigFilename, userToken)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", session)
}
