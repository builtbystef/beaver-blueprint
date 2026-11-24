// Package cli defines the Cobra commands. Each command is a thin layer over
// internal/core.
package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/builtbystef/beaver-blueprint/internal/core"
)

const workflow = `A living architectural model of this application, kept in blueprint/.

Workflow:
  - Read the blueprint before designing a change.
  - Update it in the same change as the code. Add elements with 'new' and
    follow the heuristics below.
  - Before finishing, run 'status'. For each stale element, either update it
    or confirm it still holds and stamp it with 'verify'. Then run 'check'.
  - Use the blueprint-greenfield or blueprint-brownfield skill only when the
    blueprint is first being written.

Heuristics:
  - A module is something you would assign ownership of or deploy
    separately. Never a file or a class. A typical application has five to
    fifteen top-level modules.
  - An entity is a concept the business would name, not a database table.
  - A use case is a user goal, not a screen.
  - A flow is only written when the coordination between modules is not
    obvious from the module descriptions.
  - An interface is written for every external entry point, and for internal
    boundaries only where the contract matters.
  - A decision is recorded when it would be expensive to reverse or when a
    future reader would ask "why".
  - The blueprint describes what and why. How lives in the code.`

// NewRoot builds the root command.
func NewRoot(version string) *cobra.Command {
	root := &cobra.Command{
		Use:           "beaver-blueprint",
		Short:         "A living architectural model of a software application",
		Long:          workflow,
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.AddCommand(
		newInitCmd(),
		newNewCmd(),
		newMvCmd(),
		newRmCmd(),
		newListCmd(),
		newShowCmd(),
		newCheckCmd(),
		newStatusCmd(),
		newVerifyCmd(),
		newSchemaCmd(),
		newServeCmd(),
	)
	return root
}

// Execute runs the CLI and returns the process exit code.
func Execute(version string) int {
	if err := NewRoot(version).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error: "+err.Error())
		return 1
	}
	return 0
}

// openBlueprint performs the common preamble: git checks, walk-up
// discovery, and loading.
func openBlueprint() (*core.Blueprint, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	root, err := core.Find(cwd)
	if err != nil {
		return nil, err
	}
	if err := core.RequireGit(root); err != nil {
		return nil, err
	}
	return core.Load(root)
}
