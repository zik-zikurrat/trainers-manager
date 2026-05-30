package main

import (
	"fmt"
	"trainers-manager/internal/app"
	"trainers-manager/internal/config"
)

func main() {
	// Config
	cfg := config.MustLoad()

	fmt.Printf("CONFIGS: %+v\n", cfg)

	// Run application
	if err := app.Run(cfg); err != nil {
		return
	}

}
