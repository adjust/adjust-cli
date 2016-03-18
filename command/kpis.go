package command

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/adjust/adjust-cli/adjust"
	"github.com/adjust/adjust-cli/kpis"
	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func CmdDeliverables(c *cli.Context) {
	session, err := adjust.ReadSession(adjust.DefaultConfigFilename)
	if err != nil {
		adjust.Fail("You need to be logged in first.")
	}

	params, err := kpis.NewParams(session, c)
	if err != nil {
		adjust.Fail("Failed to build params.")
	}

	req := params.NewRequest()

	fmt.Fprintf(os.Stderr, "URL: %s\n", req.URL.String())
	fmt.Fprintf(os.Stderr, "\n\n")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%#v", err)
	}
	// TODO check response code

	csvReader := csv.NewReader(res.Body)
	table, err := tablewriter.NewCSVReader(os.Stdout, csvReader, true)

	// table.SetHeader([]string{"Name", "Sign", "Rating"})
	table.SetBorder(false)

	// for _, v := range data {
	//	table.Append(v)
	// }
	table.Render() // Send output
}

func CmdCohorts(c *cli.Context) {
}
