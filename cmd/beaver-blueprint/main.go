// Command beaver-blueprint maintains a living architectural model of an
// application inside its own repository.
package main

import (
	"os"

	"github.com/builtbystef/beaver-blueprint/internal/cli"
)

// version is set by goreleaser through -ldflags.
var version = "dev"

func main() {
	os.Exit(cli.Execute(version))
}
