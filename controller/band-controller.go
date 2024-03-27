package controller

import (
	"band_api/data/request"
	"band_api/helper"
	"band_api/service"
	"encoding/json"
	"net/http"
	"strconv"

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

// @Summary creates a band
// @Description creates a band
// @Tags Bands
// @Param band body request.CreateBandRequest true "Create band"
// @Router /bands [post]
func (controller *BandController) Create(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info().Msg("Create Band")
	createBandRequest := request.CreateBandRequest{}
	json.NewDecoder(r.Body).Decode(&createBandRequest)
	controller.bandService.Create(createBandRequest)
	w.Header().Add("Content-Type", "application/json")
}

// @Summary returns a list of all bands
// @Description returns a list of all bands
// @Tags Bands
// @Router /bands [get]
func (controller *BandController) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info().Msg("FindAll bands")
	bandResponse := controller.bandService.FindAll()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bandResponse)
}

// @Summary returns a band by its id
// @Description returns a band by its id
// @Tags Bands
// @Param bandId path string true "find band by id"
// @Router /bands/{bandId} [get]
func (controller *BandController) FindById(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info().Msg("FindById band")
	bandId := r.PathValue("id")
	id, err := strconv.Atoi(bandId)
	helper.ErrorPanic(err)

	bandResponse := controller.bandService.FindById(id)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bandResponse)
}
