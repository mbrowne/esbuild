package js_parser

import (
	"testing"

	"github.com/evanw/esbuild/internal/config"
)

func expectParseErrorDCI(t *testing.T, contents string, expected string) {
	t.Helper()
	expectParseErrorCommon(t, contents, expected, config.Options{
		DCI: true,
	})
}

func expectPrintedDCI(t *testing.T, contents string, expected string) {
	t.Helper()
	expectPrintedCommon(t, contents, expected, config.Options{
		DCI: true,
	})
}

func TestDCIContextDecl(t *testing.T) {
	expectPrintedDCI(t, "\ncontext MyContext {}", "context MyContext\n")
}
