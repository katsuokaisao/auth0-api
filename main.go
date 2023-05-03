package main

import (
	"log"

	"github.com/katsuokaisao/auth0-api/api"
	"github.com/katsuokaisao/auth0-api/util"
)

func main() {
	if err := util.LoadEnv(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	api.Init()
}
