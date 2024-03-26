package main

import (
	"band_api/internal/api"
	"log"
)

func main() {
	server := api.NewAPIServer(":8087")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
