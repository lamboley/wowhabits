package main

import (
	"os"

	"github.com/lamboley/wowhabits/cmd/wowhabits/app"
)

func main() {
	if err := app.NewWowHabitsCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
