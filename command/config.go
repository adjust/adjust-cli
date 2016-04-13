package command

import (
	"github.com/adjust/adjust-cli/adjust"
	"github.com/codegangsta/cli"
)

func CmdConfig(context *cli.Context) {
	configFilename := adjust.NewSettings().ConfigFilename

	adjust.NewConfig(context).WriteConfig(configFilename)
}
