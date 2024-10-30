package main

import (
	"fmt"
	"github.com/vinicius457/BancoDeDados2/config"
	"github.com/vinicius457/BancoDeDados2/internal/router"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	router.Inicialize()
}