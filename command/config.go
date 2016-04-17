package command

import (
	"bufio"
	"os"

	"github.com/adjust/adjust-cli/adjust"
	"github.com/codegangsta/cli"
)

func CmdConfig(context *cli.Context) {
	configFilename := adjust.NewSettings().ConfigFilename

	var conf *adjust.Config

	if context.NumFlags() > 0 {
		conf = adjust.NewConfig(context).WriteConfig(configFilename)
	} else {
		conf = adjust.ReadConfig(configFilename)
	}

	buf := bufio.NewWriter(os.Stdout)
	adjust.PrintConfig(buf, conf)
	buf.Flush()

	adjust.Success()
}
