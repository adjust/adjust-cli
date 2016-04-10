package adjust

import (
	"fmt"
	"os"
)

const URLScheme = "https"
const URLHost = "api.adjust.com"

var DefaultConfigFilename = fmt.Sprintf("%s/.adjustrc", os.Getenv("HOME"))

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

func Print(message string) {
	fmt.Fprintf(os.Stdout, "%s", message)
}
