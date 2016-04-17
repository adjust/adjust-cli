package adjust

import (
	"bufio"
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

func PrintConfig(buf *bufio.Writer, conf *Config) {
	printConfig(buf, conf)
}

func printConfig(buf *bufio.Writer, conf *Config) {
	var err error

	if conf.UserToken != "" {
		_, err = fmt.Fprintf(buf, "user_token: %s\n", conf.UserToken)
	}

	if conf.AppToken != "" {
		_, err = fmt.Fprintf(buf, "app_token: %s\n", conf.AppToken)
	}

	if conf.AppTokens != nil {
		_, err = fmt.Fprintf(buf, "app_tokens: %s\n", strings.Join(conf.AppTokens, ","))
	}

	if err != nil {
		Fail(err.Error())
	}
}
