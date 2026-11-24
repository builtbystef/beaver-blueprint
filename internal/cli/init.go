package cli

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/builtbystef/beaver-blueprint/internal/core"
)

func newInitCmd() *cobra.Command {
	var agent string
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Scaffold blueprint/, write the AGENTS.md pointer, install skills",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}
			if err := core.RequireGit(cwd); err != nil {
				return err
			}
			// TODO: AGENTS.md pointer and skill installation (--agent).
			return core.Init(cwd)
		},
	}
	cmd.Flags().StringVar(&agent, "agent", "", "install skills for this agent only")
	return cmd
}
