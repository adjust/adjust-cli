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
	UserToken   string
}

func NewParams(session *adjust.Session, context *cli.Context) (*Params, error) {
	res := &Params{
		UserToken:   session.UserToken,
		AppTokens:   strings.Split(context.String("app_tokens"), ","),
		Trackers:    strings.Split(context.String("trackers"), ","),
		KPIs:        strings.Split(context.String("kpis"), ","),
		OSNames:     strings.Split(context.String("os_names"), ","),
		Countries:   strings.Split(context.String("countries"), ","),
		DeviceTypes: strings.Split(context.String("device_types"), ","),
		Grouping:    strings.Split(context.String("grouping"), ","),
	}

	return res, nil
}

func (params *Params) NewRequest(endpoint string) *http.Request {
	req, err := http.NewRequest("GET", urlFromParams(endpoint, params).String(), nil)
	if err != nil {
		adjust.Fail("Failed to build URL.")
	}

	header := adjust.DefaultHeaders(params.UserToken)
	header.Add("Accept", "text/csv")
	req.Header = *header

	return req
}

func urlFromParams(endpoint string, params *Params) *url.URL {
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

	if appendTrackerFilters {
		appendParam(q, "tracker_filter", params.Trackers)
	}

	return &url.URL{
		Scheme:   adjust.URLScheme,
		Host:     adjust.URLHost,
		Path:     path,
		RawQuery: q.Encode(),
	}
}

func appendParam(q url.Values, name string, values []string) {
	if len(values) > 0 && values[0] != "" {
		q.Add(name, strings.Join(values, ","))
	}
}
