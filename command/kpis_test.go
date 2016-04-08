package command

import (
	"fmt"
	"testing"
)

func fail(t *testing.T, headline string, actual string, expected string) {
	t.Error(headline, fmt.Sprintf("\nActual:   %s\n", actual), fmt.Sprintf("\nExpected: %s\n", expected))
}

func TestHeader(t *testing.T) {
	// 1. test verbose output
	// 2. test csv save to tmp file
	// 3. test output on HTTP error
}
