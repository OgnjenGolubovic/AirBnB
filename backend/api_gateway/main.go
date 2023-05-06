package main

import (
	"github.com/OgnjenGolubovic/AirBnB/backend/api_gateway/startup"
	"github.com/OgnjenGolubovic/AirBnB/backend/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
