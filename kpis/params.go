package kpis

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/adjust/adjust-cli/adjust"
	"github.com/codegangsta/cli"
)

const URLPath = "kpis/v1"

type Params struct {
	Trackers    []string
	AppTokens   []string
	KPIs        []string
	OSNames     []string
	Countries   []string
	DeviceTypes []string
	Grouping    []string
	Start       string
	End         string
	UserToken   string
}

func NewParams(endpoint string, config *adjust.Config, context *cli.Context) (*Params, error) {
	res := &Params{
		UserToken:   config.UserToken,
		AppTokens:   appTokens(endpoint, config, adjust.CommaSeparatedParam(context, "app_tokens")),
		Trackers:    commaSeparatedParam(context, "trackers"),
		KPIs:        commaSeparatedParam(context, "kpis"),
		OSNames:     commaSeparatedParam(context, "os_names"),
		Countries:   commaSeparatedParam(context, "countries"),
		DeviceTypes: commaSeparatedParam(context, "device_types"),
		Start:       context.String("start"),
		End:         context.String("end"),
		Grouping:    commaSeparatedParam(context, "grouping"),
	}

	return res, nil
}

func appTokens(endpoint string, config *adjust.Config, contextTokens []string) []string {
	if contextTokens != nil {
		return contextTokens
	}

	if config.AppTokens != nil && endpoint == "deliverables" {
		return config.AppTokens
	}

	if config.AppToken != "" {
		return []string{config.AppToken}
	}

	return nil
}

func (params *Params) NewRequest(endpoint string, scheme string, host string) *http.Request {
	req, err := http.NewRequest("GET", buildURL(endpoint, scheme, host, params).String(), nil)
	if err != nil {
		adjust.Fail("Failed to build URL.")
	}

	header := adjust.DefaultHeaders(params.UserToken)
	header.Add("Accept", "text/csv")
	req.Header = *header

	return req
}

func buildURL(endpoint string, scheme string, host string, params *Params) *url.URL {
	path := URLPath
	appendTrackerFilters := true
	q := url.Values{}

	if len(params.AppTokens) == 1 && params.AppTokens[0] != "" {
		path = fmt.Sprintf("%s/%s", path, params.AppTokens[0])

		if len(params.Trackers) == 1 && params.Trackers[0] != "" {
			path = fmt.Sprintf("%s/trackers/%s", path, params.Trackers[0])
			appendTrackerFilters = false
		}
	} else {
		q.Add("app_tokens", strings.Join(params.AppTokens, ","))
	}

	if endpoint == "cohorts" {
		path = fmt.Sprintf("%s/cohorts", path)
	}
	path = fmt.Sprintf("%s.csv", path)

	appendParam(q, "kpis", params.KPIs)
	appendParam(q, "grouping", params.Grouping)
	appendParam(q, "os_names", params.OSNames)
	appendParam(q, "countries", params.Countries)
	appendParam(q, "device_types", params.DeviceTypes)
	addParam(q, "start_date", params.Start)
	addParam(q, "end_date", params.End)

	if appendTrackerFilters {
		appendParam(q, "tracker_filter", params.Trackers)
	}

	return &url.URL{
		Scheme:   scheme,
		Host:     host,
		Path:     path,
		RawQuery: q.Encode(),
	}
}

func appendParam(q url.Values, name string, values []string) {
	if len(values) > 0 && values[0] != "" {
		q.Add(name, strings.Join(values, ","))
	}
}

func addParam(q url.Values, name string, value string) {
	if value != "" {
		q.Add(name, value)
	}
}

func commaSeparatedParam(context *cli.Context, paramName string) []string {
	if context.String(paramName) == "" {
		return nil
	}

	return strings.Split(context.String(paramName), ",")
}
