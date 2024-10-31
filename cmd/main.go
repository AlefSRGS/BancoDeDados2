package main

import (
	"github.com/vinicius457/BancoDeDados2/config"
	"github.com/vinicius457/BancoDeDados2/internal/router"
)

var(
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Inicialize()
}