package cli

import (
	"github.com/spf13/cobra"

	"github.com/builtbystef/beaver-blueprint/internal/core"
)

func newCheckCmd() *cobra.Command {
	var opts core.CheckOptions
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Validate the blueprint",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			bp, err := openBlueprint()
			if err != nil {
				return err
			}
			_, err = bp.Check(opts)
			// TODO: print diagnostics, exit 1 on any error.
			return err
		},
	}
	cmd.Flags().BoolVar(&opts.Strict, "strict", false, "treat warnings as errors")
	cmd.Flags().BoolVar(&opts.Fix, "fix", false, "apply mechanical repairs")
	return cmd
}
