package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/builtbystef/beaver-blueprint/internal/core"
)

func newNewCmd() *cobra.Command {
	var title string
	cmd := &cobra.Command{
		Use:   "new <type> <id>",
		Short: "Create an element with correct frontmatter and headings",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			spec, ok := core.LookupType(args[0])
			if !ok {
				return fmt.Errorf("unknown type %q", args[0])
			}
			if !core.ValidID(args[1]) {
				return fmt.Errorf("invalid id %q", args[1])
			}
			bp, err := openBlueprint()
			if err != nil {
				return err
			}
			_, err = bp.NewElement(spec.Type, args[1], title)
			return err
		},
	}
	cmd.Flags().StringVar(&title, "title", "", "display title (defaults from the id)")
	return cmd
}
