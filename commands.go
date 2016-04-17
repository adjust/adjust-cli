package main

import (
	"fmt"

	"github.com/adjust/adjust-cli/adjust"
	"github.com/adjust/adjust-cli/command"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{}

var ConfigFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "user-token, u",
		Usage: "Your adjust user token, available from your Dashboard.",
	},
	cli.StringFlag{
		Name:  "app-tokens",
		Usage: "A comma-separated list of your adjust app tokens.",
	},
	cli.StringFlag{
		Name:  "app-token",
		Usage: "Setup an app token for your requests. If `--app-tokens` is also given, that takes precedence for all requests that support mutiple apps",
	},
}

var KPIsFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "app-tokens, a",
		Usage: "The app token for the request",
	},
	cli.StringFlag{
		Name:  "trackers, t",
		Usage: "If a tracker token is given, tracker filtering will be applied on the request",
	},
	cli.StringFlag{
		Name:  "platforms, p",
		Usage: "OS Name filtering. Example: `--platforms ios,android`",
	},
	cli.StringFlag{
		Name:  "countries, c",
		Usage: "Country filtering using ISO Alpha 2 country codes. Example: `--countries us,fr`",
	},
	cli.StringFlag{
		Name:  "devices, d",
		Usage: "Device Type filtering. Example: `--devices phone,tablet`",
	},
	cli.StringFlag{
		Name:  "group, g",
		Value: "network",
		Usage: "Grouping for the data. Example: `--group apps,networks,campaigns`",
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
	cli.BoolFlag{
		Name:  "sandbox",
		Usage: "Return only sandbox data.",
	},
	cli.BoolFlag{
		Name:  "url-only",
		Usage: "Do not query the API, instead only print generated URL for the request.",
	},
}

var DeliverablesKPIsFlag = []cli.Flag{
	cli.StringFlag{
		Name:  "kpis, k",
		Value: "installs,clicks,sessions",
		Usage: "The KPIs for the request.",
	},
}

var CohortsKPIsFlag = []cli.Flag{
	cli.StringFlag{
		Name:  "kpis, k",
		Value: "retained_users",
		Usage: "The KPIs for the request.",
	},
	cli.StringFlag{
		Name:  "period",
		Usage: "Period for the cohort report - day/week/month",
	},
	cli.StringFlag{
		Name:  "events, e",
		Usage: "Comma-separated list of event tokens for event-based cohorts",
	},
}

var Commands = []cli.Command{
	{
		Name:   "config",
		Usage:  "configure user-token and app-tokens options to adjust",
		Action: command.CmdConfig,
		Flags:  ConfigFlags,
	},
	{
		Name:        "deliverables",
		Aliases:     []string{"d"},
		Usage:       "deliverables data for your apps.",
		Action:      command.CmdDeliverables,
		Description: "Deliverable KPIs are installs, clicks, sessions, DAU, events and so on. Check https://docs.adjust.com/en/kpi-service/",
		Flags:       append(KPIsFlags, DeliverablesKPIsFlag...),
	},
	{
		Name:        "cohorts",
		Aliases:     []string{"c"},
		Usage:       "cohorts data for your apps.",
		Action:      command.CmdCohorts,
		Description: "Cohort KPIs are Retained Users, Lifetime Value and so on. Check https://docs.adjust.com/en/kpi-service/",
		Flags:       append(KPIsFlags, CohortsKPIsFlag...),
	},
}

func CommandNotFound(c *cli.Context, command string) {
	adjust.Fail(fmt.Sprintf("%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name))
}
