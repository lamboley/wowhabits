package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/lamboley/wowhabits/pkg/api"
)

var version = "dev"

func main() {
	var addr string
	serve := &cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP API server",
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
			defer stop()

			srv := &http.Server{Addr: addr, Handler: api.NewRouter(), ReadHeaderTimeout: 5 * time.Second}

			shutdown := make(chan error, 1)
			go func() {
				<-ctx.Done()
				stop()
				shutdown <- srv.Shutdown(context.Background())
			}()

			cmd.PrintErrf("listening on %s\n", addr)
			if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				return err
			}
			return <-shutdown
		},
	}
	serve.Flags().StringVar(&addr, "addr", ":8080", "address to listen on")

	root := &cobra.Command{Use: "wowhabits", Short: "WoW Habits server", Version: version, SilenceUsage: true}
	root.AddCommand(serve)

	if root.Execute() != nil {
		os.Exit(1)
	}
}
