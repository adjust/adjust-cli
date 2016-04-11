package command

import (
	"fmt"
	"log"

	"github.com/adjust/adjust-cli/adjust"
	"github.com/codegangsta/cli"
)

func CmdConfig(c *cli.Context) {
	if c.String("user_token") == "" {
		adjust.Fail(fmt.Sprintf("You need to provide a user token. See '%s config --help'", c.App.Name))
	}

	userToken := c.String("user_token")
	session, err := adjust.NewSession(adjust.DefaultConfigFilename, userToken)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", session)
}
