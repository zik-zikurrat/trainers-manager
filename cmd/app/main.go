package main

import (
	"trainers-manager/internal/app"
	"trainers-manager/internal/config"
)

func main() {
	// Config
	cfg := config.MustLoad()

	// Run application
	if err := app.Run(cfg); err != nil {
		return
	}

}
