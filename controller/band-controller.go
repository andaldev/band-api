package controller

import (
	"band_api/service"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type BandController struct {
	bandService service.BandService
}

func NewBandController(service service.BandService) *BandController {
	return &BandController{
		bandService: service,
	}
}

func (controller *BandController) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info().Msg("FindAll bands")
	bandResponse := controller.bandService.FindAll()
	json.NewEncoder(w).Encode(bandResponse)

}
