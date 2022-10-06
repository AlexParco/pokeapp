package main

import (
	"github.com/alexparco/pokeapp-api/app"
	"github.com/alexparco/pokeapp-api/config"
)

func main() {
	cfg := config.ReadConfig("./config.yaml")

	api := app.NewApi(cfg)

	api.Run()
}
