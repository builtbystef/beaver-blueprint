package cli

import "github.com/spf13/cobra"

func newVerifyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "verify <id>...",
		Short: "Stamp elements as reviewed at HEAD",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			bp, err := openBlueprint()
			if err != nil {
				return err
			}
			return bp.Verify(args...)
		},
	}
}
