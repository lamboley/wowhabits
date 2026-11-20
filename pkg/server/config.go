package server

import "github.com/spf13/pflag"

// RegisterFlags registers the server flags on fs, with their default values.
func (cfg *Config) RegisterFlags(fs *pflag.FlagSet) {
	fs.StringVar(&cfg.Addr, "addr", ":8080", "The address the HTTP server listens on.")
}
