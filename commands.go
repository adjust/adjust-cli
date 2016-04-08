package main

import (
	"fmt"

	"github.com/adjust/adjust-cli/adjust"
	"github.com/adjust/adjust-cli/command"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{}

var KPIsFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "kpis, k",
		Value: "installs,clicks,sessions",
		Usage: "The KPIs for the request.",
	},
	cli.StringFlag{
		Name:  "app_tokens, a",
		Usage: "The app token for the request",
	},
	cli.StringFlag{
		Name:  "grouping, g",
		Value: "network,campaign",
		Usage: "Grouping for the data, e.g. apps,networks,campaigns",
	},
	cli.BoolFlag{
		Name:  "verbose",
		Usage: "Print additional information when running the requests.",
	},
}

var Commands = []cli.Command{
	{
		Name:   "login",
		Usage:  "persists the user token in a config file.",
		Action: command.CmdLogin,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "logout",
		Usage:  "erase all user data from config file.",
		Action: command.CmdLogout,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "deliverables",
		Usage:  "installs, clicks, sessions, events etc. data.",
		Action: command.CmdDeliverables,
		Flags:  KPIsFlags,
	},
	{
		Name:   "cohorts",
		Usage:  "get cohorts data.",
		Action: command.CmdCohorts,
		Flags:  KPIsFlags,
	},
}

func DefaultAction(c *cli.Context) {
	command.CmdDeliverables(c)
}

func CommandNotFound(c *cli.Context, command string) {
	adjust.Fail(fmt.Sprintf("%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name))
}
