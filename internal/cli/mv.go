package cli

import "github.com/spf13/cobra"

func newMvCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mv <old-id> <new-id>",
		Short: "Rename an element and rewrite every reference to it",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			bp, err := openBlueprint()
			if err != nil {
				return err
			}
			return bp.Move(args[0], args[1])
		},
	}
}
