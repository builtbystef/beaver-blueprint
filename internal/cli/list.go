package cli

import (
	"github.com/spf13/cobra"

	"github.com/builtbystef/beaver-blueprint/internal/core"
)

func newListCmd() *cobra.Command {
	var typ string
	cmd := &cobra.Command{
		Use:   "list",
		Short: "One line per element: type, id, title, stale marker",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := openBlueprint(); err != nil {
				return err
			}
			return core.ErrNotImplemented
		},
	}
	cmd.Flags().StringVar(&typ, "type", "", "only list elements of this type")
	return cmd
}
