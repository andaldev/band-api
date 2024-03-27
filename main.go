package main

import (
	"band_api/controller"
	"band_api/db"
	"band_api/repository"
	"band_api/router"
	"band_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// import (
// 	"band_api/internal/api"
// 	"log"
// )

// // @title Swagger band API
// // @version 1.0
// // @description This is a sample server band server.
// // @termsOfService http://swagger.io/terms/
// func main() {
// 	server := api.NewAPIServer(":8080")
// 	if err := server.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {

	log.Info().Msg("Started server!")
	db := db.InitializeDatabase()
	validate := validator.New()

	bandRepository := repository.NewBandRepositoryImpl(db)
	bandService := service.NewBandServiceImpl(bandRepository, validate)
	bandController := controller.NewBandController(bandService)

	routes := router.NewRouter(bandController)

	server := &http.Server{
		Addr: ":8090",
	}

	http.ListenAndServe(server.Addr, routes)
}
