package cli

import (
	"github.com/spf13/cobra"

	"github.com/builtbystef/beaver-blueprint/internal/core"
)

func newSchemaCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "schema [<type>]",
		Short: "Print the frontmatter JSON Schema",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// No git or blueprint required. TODO: build the document with
			// server.NewAPI(http.NewServeMux(), nil).OpenAPI() and print the
			// frontmatter component schemas, so it can never go stale.
			return core.ErrNotImplemented
		},
	}
}
