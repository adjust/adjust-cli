package main

import "github.com/codegangsta/cli"

const Name string = "adjust"
const Version string = "0.0.1"

func configureCLI() {
	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .Flags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{end}}
   {{if .Version}}
VERSION:
   {{.Version}}
   {{end}}{{if len .Authors}}
AUTHOR(S):
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
   {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}
`

	cli.CommandHelpTemplate = `NAME:
   {{.HelpName}} - {{.Usage}}

USAGE:
   {{.HelpName}}{{if .Flags}} [options]{{end}} {{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if .Flags}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}
`
}
