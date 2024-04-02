package main

import (
	"band_api/api"
)

// @title 	Band Service API
// @version	1.0
// @description A Band service API in Go
func main() {

	server := api.NewAPIServer(":8080")
	server.Run()
}
