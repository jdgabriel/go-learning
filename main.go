package main

import (
	"github.com/jdgabriel/go-learning/config"
	"github.com/jdgabriel/go-learning/router"
)

var logger config.Logger

func main() {
	logger = *config.GetLogger("API")
	// Init configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error", err)
		return
	}
	// Init router package
	router.Initialize()
}
