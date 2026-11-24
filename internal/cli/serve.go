package cli

import (
	"github.com/spf13/cobra"

	"github.com/builtbystef/beaver-blueprint/internal/server"
	"github.com/builtbystef/beaver-blueprint/web"
)

func newServeCmd() *cobra.Command {
	var opts server.Options
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the local web UI",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			bp, err := openBlueprint()
			if err != nil {
				return err
			}
			opts.Static = web.Dist()
			return server.Run(cmd.Context(), bp, opts)
		},
	}
	cmd.Flags().IntVar(&opts.Port, "port", server.DefaultPort, "port to listen on")
	cmd.Flags().BoolVar(&opts.Open, "open", false, "open the browser")
	return cmd
}
