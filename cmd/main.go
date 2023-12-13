package main

import (
	"log"
	"taskTracker/config"
	"taskTracker/internal/server"
	"taskTracker/pkg/pgConnector"
)

func main() {
	cfg, err := config.LoadConfig(config.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	conn := pgConnector.Connect()

	servers := server.NewServer(cfg, conn)

	servers.Run()
}
