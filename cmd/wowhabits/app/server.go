// Package app does all of the work necessary to configure and run a WowHabits app process.
package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/lamboley/wowhabits/pkg/server"
)

// NewWowHabitsCommand creates a *cobra.Command object
func NewWowHabitsCommand() *cobra.Command {
	cfg := NewConfig()

	cmd := &cobra.Command{
		Use:          "wowhabits",
		Short:        "Run the WoW Habits server",
		Version:      gitVersion,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cfg.Run(cmd.Context())
		},
	}

	fs := cmd.Flags()
	cfg.AddAllFlags(fs)

	return cmd
}

// Run builds every component from the config
func (c *Config) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv, err := server.New(c.Server)
	if err != nil {
		return err
	}

	shutdown := make(chan error, 1)
	go func() {
		<-ctx.Done()
		stop()
		shutdown <- srv.Stop(context.Background())
	}()

	if err := srv.Start(); err != nil {
		return err
	}

	return <-shutdown
}
