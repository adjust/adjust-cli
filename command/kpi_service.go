package command

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/adjust/adjust-cli/adjust"
	"github.com/adjust/adjust-cli/kpis"
	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func CmdDeliverables(c *cli.Context) {
	performKPIServiceRequest("deliverables", c)
}

func CmdCohorts(c *cli.Context) {
	performKPIServiceRequest("cohorts", c)
}

func performKPIServiceRequest(endpoint string, c *cli.Context) {
	validateContext(c)

	session, err := adjust.ReadSession(adjust.DefaultConfigFilename)
	if err != nil {
		adjust.Fail("You need to be logged in first.")
	}

	params, err := kpis.NewParams(session, c)
	if err != nil {
		adjust.Fail("Failed to build params.")
	}

	req := params.NewRequest(endpoint)

	unescaped, err := url.QueryUnescape(req.URL.String())
	if err != nil {
		adjust.Fail("Failed to unescape url.")
	}

	verbose := c.Bool("verbose")
	if verbose {
		adjust.Notify(fmt.Sprintf("Requesting URL: %s\n", unescaped))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if verbose {
			adjust.Fail(err.Error())
		} else {
			adjust.Fail("Could not connect to the adjust API.")
		}
	}
	defer res.Body.Close()

	if verbose {
		handleResponseVerbose(res)
	} else {
		handleResponse(res)
	}

	if c.Bool("csv") || c.String("file") != "" {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			adjust.Fail("Could not read CSV response from KPI Service.")
		}

		w := os.Stdout
		if c.String("file") != "" {
			w, err = os.OpenFile(c.String("file"), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
			if err != nil {
				adjust.Fail(err.Error())
			}
		}

		buf := bufio.NewWriter(w)
		_, err = buf.Write(body)
		if err != nil {
			adjust.Fail(err.Error())
		}
		buf.Flush()

		adjust.Success()
	}

	csvReader := csv.NewReader(res.Body)
	table, err := tablewriter.NewCSVReader(os.Stdout, csvReader, true)

	// table.SetHeader([]string{"Name", "Sign", "Rating"})
	table.SetBorder(false)

	// for _, v := range data {
	//	table.Append(v)
	// }
	table.Render() // Send output
}

func handleResponse(res *http.Response) {
	if res.StatusCode >= 500 {
		adjust.Fail("KPI Service error encountered.")
	}

	if res.StatusCode >= 400 {
		text, err := ioutil.ReadAll(res.Body)
		if err != nil {
			adjust.Fail("Could not parse KPI Service response.")
		}
		adjust.Fail(string(text))
	}

	if res.StatusCode != 200 {
		adjust.Fail("An error occurred connecting to adjust KPI service.")
	}
}

func handleResponseVerbose(res *http.Response) {
	if res.StatusCode >= 500 {
		adjust.Fail(fmt.Sprintf("KPI Service error %s encountered", res.Status))
	}

	if res.StatusCode >= 400 {
		adjust.Error(fmt.Sprintf("Incorrect request to the KPI Service: %s", res.Status))
		text, err := ioutil.ReadAll(res.Body)
		if err != nil {
			adjust.Fail("Could not parse KPI Service response.")
		}
		adjust.Fail(string(text))
	}

	if res.StatusCode != 200 {
		adjust.Fail(fmt.Sprintf("Unexpected HTTP response from the KPI Service: %s", res.Status))
	}
}

func validateContext(c *cli.Context) {
	if c.String("start") == "" && c.String("end") != "" || c.String("start") != "" && c.String("end") == "" {
		adjust.Fail("--start and --end dates must be both given.")
	}
}
