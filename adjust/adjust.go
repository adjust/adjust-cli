package adjust

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

func Success() {
	os.Exit(0)
}

func Fail(message string) {
	Error(message)
	os.Exit(2)
}

func Error(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", message)
}

func Notify(message string) {
	fmt.Fprintf(os.Stdout, "%s\n", message)
}

func DefaultHeaders(userToken string) *http.Header {
	res := http.Header{}
	res.Add("Authorization", fmt.Sprintf("Token token=%s", userToken))
	res.Add("X-Adjust-CLI", "1")

	return &res
}

func CommaSeparatedParam(context *cli.Context, paramName string) []string {
	return commaSeparatedParam(context, paramName)
}

func commaSeparatedParam(context *cli.Context, paramName string) []string {
	if context.String(paramName) == "" {
		return nil
	}

	return strings.Split(context.String(paramName), ",")
}
