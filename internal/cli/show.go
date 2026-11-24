package cli

import (
	"github.com/spf13/cobra"

	"github.com/builtbystef/beaver-blueprint/internal/core"
)

func newShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show <id>",
		Short: "Print the file, then resolved relationships in both directions",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := openBlueprint(); err != nil {
				return err
			}
			return core.ErrNotImplemented
		},
	}
}
