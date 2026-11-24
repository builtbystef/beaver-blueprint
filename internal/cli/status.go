package cli

import "github.com/spf13/cobra"

func newStatusCmd() *cobra.Command {
	var committed bool
	cmd := &cobra.Command{
		Use:   "status",
		Short: "List stale elements and the files that changed",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			bp, err := openBlueprint()
			if err != nil {
				return err
			}
			_, err = bp.Drift()
			// TODO: print committed and working-tree groups.
			return err
		},
	}
	cmd.Flags().BoolVar(&committed, "committed", false, "hide working-tree drift")
	return cmd
}
