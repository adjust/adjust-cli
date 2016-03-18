package kpis

import (
	"fmt"
	"testing"
)

func fail(t *testing.T, headline string, actual string, expected string) {
	t.Error(headline, fmt.Sprintf("\nActual:   %s\n", actual), fmt.Sprintf("\nExpected: %s\n", expected))
}

type URLTestCase struct {
	Params   Params
	Expected string
}

func TestURL(t *testing.T) {
	cases := []URLTestCase{
		{
			Params{
				AppTokens: []string{"abcdef"},
				KPIs:      []string{"installs", "sessions"},
			},
			"https://api.adjust.com/kpis/v1/abcdef.csv?kpis=installs%2Csessions",
		},
		{
			Params{
				AppTokens: []string{"abcdef"},
				Trackers:  []string{"my-token"},
				KPIs:      []string{"installs"},
			},
			"https://api.adjust.com/kpis/v1/abcdef/trackers/my-token.csv?kpis=installs",
		},
		{
			Params{
				AppTokens: []string{"abcdef", "123456"},
				KPIs:      []string{"installs"},
			},
			"https://api.adjust.com/kpis/v1.csv?app_tokens=abcdef%2C123456&kpis=installs",
		},
		{
			Params{
				AppTokens:   []string{"abcdef", "123456"},
				Trackers:    []string{"my-tracker"},
				KPIs:        []string{"installs"},
				OSNames:     []string{"ios"},
				Countries:   []string{"de", "us"},
				DeviceTypes: []string{"phone"},
				Grouping:    []string{"apps"},
			},
			"https://api.adjust.com/kpis/v1.csv?app_tokens=abcdef%2C123456&countries=de%2Cus&device_types=phone&grouping=apps&kpis=installs&os_names=ios&tracker_filter=my-tracker",
		},
	}

	for _, tc := range cases {
		actual := tc.Params.NewRequest().URL.String()

		if actual != tc.Expected {
			fail(t, "URL test failed:", actual, tc.Expected)
		}
	}
}

func TestHeader(t *testing.T) {
	params := Params{UserToken: "my-user-token"}
	req := params.NewRequest()

	acceptHeader := "text/csv"
	if v := req.Header.Get("Accept"); v != acceptHeader {
		fail(t, "Header test failed: ", v, acceptHeader)
	}

	acceptHeader = "Token token=my-user-token"
	if v := req.Header.Get("Authorization"); v != acceptHeader {
		fail(t, "Header test failed: ", v, acceptHeader)
	}
}
