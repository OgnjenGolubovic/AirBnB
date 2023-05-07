package main

import (
	"reservation_service/startup"
	cfg "reservation_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
