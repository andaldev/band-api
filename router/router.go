package router

import (
	"band_api/controller"
	_ "band_api/docs"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewRouter(bandController *controller.BandController) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	router.HandleFunc("GET /bands", bandController.FindAll)
	router.HandleFunc("GET /bands/{id}", bandController.FindById)

	return router
}
