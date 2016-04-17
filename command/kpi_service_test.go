package command

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func fail(t *testing.T, headline string, actual string, expected string) {
	t.Error(headline, fmt.Sprintf("\nActual:   %s\n", actual), fmt.Sprintf("\nExpected: %s\n", expected))
}

func compile(t *testing.T) {
	err := exec.Command("sh", "-c", "go build -o adjust-cli-test github.com/adjust/adjust-cli").Run()
	if err != nil {
		t.Error(err.Error())
	}
}

func cleanCompiled() {
	_ = os.Remove("adjust-cli-test")
}

func testCommand(t *testing.T, args string) string {
	out, err := exec.Command("sh", "-c", fmt.Sprintf("./adjust-cli-test %s --url-only", args)).Output()
	if err != nil {
		t.Error(err)
	}

	return string(out)
}

type TestCase struct {
	Args        string
	ExpectedURL string
}

func TestKPICommands(t *testing.T) {
	compile(t)
	defer cleanCompiled()

	cases := []TestCase{
		TestCase{
			"deliverables -a abc",
			"https://api.adjust.com/kpis/v1/abc.csv?grouping=network&kpis=installs,clicks,sessions",
		},
		TestCase{
			"d -a abc",
			"https://api.adjust.com/kpis/v1/abc.csv?grouping=network&kpis=installs,clicks,sessions",
		},
		TestCase{
			"deliverables -a abc -k installs --platforms iphone --countries de,us --sandbox",
			"https://api.adjust.com/kpis/v1/abc.csv?countries=de,us&grouping=network&kpis=installs&os_names=iphone&sandbox=true",
		},
		TestCase{
			"deliverables -a abc,def --group apps --kpis installs,clicks",
			"https://api.adjust.com/kpis/v1.csv?app_tokens=abc,def&grouping=apps&kpis=installs,clicks",
		},
		TestCase{
			"cohorts -a abc",
			"https://api.adjust.com/kpis/v1/abc/cohorts.csv?grouping=network&kpis=retained_users",
		},
		TestCase{
			"c -a abc",
			"https://api.adjust.com/kpis/v1/abc/cohorts.csv?grouping=network&kpis=retained_users",
		},
	}

	for _, tc := range cases {
		output := testCommand(t, tc.Args)

		if output != fmt.Sprintf("%s\n", tc.ExpectedURL) {
			fail(t, "Invalid URL", output, tc.ExpectedURL)
		}
	}
}
