package cli

import "github.com/spf13/cobra"

func newRmCmd() *cobra.Command {
	var force bool
	cmd := &cobra.Command{
		Use:   "rm <id>",
		Short: "Delete an element; fails if referenced unless --force",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			bp, err := openBlueprint()
			if err != nil {
				return err
			}
			return bp.Remove(args[0], force)
		},
	}
	cmd.Flags().BoolVar(&force, "force", false, "strip references to the element")
	return cmd
}
