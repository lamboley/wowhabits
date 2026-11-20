package app

import (
	"github.com/spf13/pflag"

	"github.com/lamboley/wowhabits/pkg/server"
)

// Config is the root config for WoWHabits.
type Config struct {
	Server server.Config
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	return &Config{}
}

// AddAllFlags regiser flags for each component
func (c *Config) AddAllFlags(fs *pflag.FlagSet) {
	c.Server.RegisterFlags(fs)
}
