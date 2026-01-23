package main

import (
	"kasir-api/internal/server"
	"log"
)

var version = "dev"

func main() {
	addr := ":8080"
	log.Printf("Server %s running on %s \n", addr, version)

	server.Start(addr)
}
