package main

import (
	"github.com/OgnjenGolubovic/AirBnB/backend/user_service/startup"
	cfg "github.com/OgnjenGolubovic/AirBnB/backend/user_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
