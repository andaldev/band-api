package router

import (
	"band_api/controller"
	"net/http"
)

func NewRouter(bandController *controller.BandController) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /band", bandController.FindAll)

	return router
}
