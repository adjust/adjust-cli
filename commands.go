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
		Name:  "app_tokens, a",
		Usage: "The app token for the request",
	},
	cli.StringFlag{
		Name:  "trackers, t",
		Usage: "If a tracker token is given, tracker filtering will be applied on the request",
	},
	cli.StringFlag{
		Name:  "os_names, o",
		Usage: "OS Name filtering. Example: `--os_names 'ios,android'`",
	},
	cli.StringFlag{
		Name:  "countries, c",
		Usage: "Country filtering using ISO Alpha 2 country codes. Example: `--countries 'us,fr'`",
	},
	cli.StringFlag{
		Name:  "device_types, d",
		Usage: "Device Type filtering. Example: `--device_types phone`",
	},
	cli.StringFlag{
		Name:  "grouping, g",
		Value: "network,campaign",
		Usage: "Grouping for the data. Example: `--grouping apps,networks,campaigns`",
	},
	cli.StringFlag{
		Name:  "file",
		Usage: "Save the response data to a file instead of printing it to the terminal.",
	},
	cli.BoolFlag{
		Name:  "csv",
		Usage: "Don't prettify the KPIs data, but output in CSV format.",
	},
	cli.BoolFlag{
		Name:  "verbose",
		Usage: "Print additional information when running the requests.",
	},
	cli.StringFlag{
		Name:  "start",
		Usage: "The start date of the request period in format YYYY-MM-DD. Example: `--start 2015-01-01`.",
	},
	cli.StringFlag{
		Name:  "end",
		Usage: "The end date of the request period in format YYYY-MM-DD. Example: `--end 2015-01-31`.",
	},
}

var DeliverablesKPIsFlag = cli.StringFlag{
	Name:  "kpis, k",
	Value: "installs,clicks,sessions",
	Usage: "The KPIs for the request.",
}

var CohortsKPIsFlag = cli.StringFlag{
	Name:  "kpis, k",
	Value: "retained_users",
	Usage: "The KPIs for the request.",
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
		Flags:  append(KPIsFlags, DeliverablesKPIsFlag),
	},
	{
		Name:   "cohorts",
		Usage:  "get cohorts data.",
		Action: command.CmdCohorts,
		Flags:  append(KPIsFlags, CohortsKPIsFlag),
	},
}

func CommandNotFound(c *cli.Context, command string) {
	adjust.Fail(fmt.Sprintf("%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name))
}
