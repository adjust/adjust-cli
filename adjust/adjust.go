package adjust

import (
	"fmt"
	"os"
)

const URLScheme = "https"
const URLHost = "api.adjust.com"

var DefaultConfigFilename = fmt.Sprintf("%s/.adjustrc", os.Getenv("HOME"))

type Options struct {
}

func Fail(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", message)
	os.Exit(2)
}
