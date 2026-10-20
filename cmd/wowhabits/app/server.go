// Package app does all of the work necessary to configure and run a WowHabits app process.
package app

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lamboley/wowhabits/pkg/version"
)

// NewWowHabitsCommand creates a *cobra.Command object
func NewWowHabitsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "wowhabits",
		Short:        "Run the WoW Habits API server",
		Version:      version.Get(),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Print("This is the root command")

			return nil
		},
	}

	return cmd
}
