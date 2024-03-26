package main

import (
	"band_api/internal/api"
	"log"
)

// @title Swagger band API
// @version 1.0
// @description This is a sample server band server.
// @termsOfService http://swagger.io/terms/
func main() {
	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
