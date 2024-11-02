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

	err := config.DBConnection()
	if err != nil {
        logger.Errorf("Error connecting to database: %v", err)
		return
    }

	router.Inicialize()
}