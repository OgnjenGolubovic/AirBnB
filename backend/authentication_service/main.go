package main

import (
	"authentication_service/startup"
	cfg "authentication_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
