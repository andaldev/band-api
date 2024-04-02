package controller

import (
	"band_api/data/request"
	"band_api/helper"
	"band_api/service"
	"encoding/json"
	"net/http"
	"strconv"
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
	createBandRequest := request.CreateBandRequest{}
	json.NewDecoder(r.Body).Decode(&createBandRequest)
	controller.bandService.Create(createBandRequest)
	w.Header().Add("Content-Type", "application/json")
}

// @Summary updates a band
// @Description updates a band
// @Tags Bands
// @Param bandId path string true "update band by id"
// @Param band body request.CreateBandRequest true "Update band"
// @Router /bands/{bandId} [patch]
func (controller *BandController) Update(w http.ResponseWriter, r *http.Request) {
	updateBandRequest := request.UpdateBandRequest{}
	err := json.NewDecoder(r.Body).Decode(&updateBandRequest)
	helper.ErrorPanic(err)

	bandId := r.PathValue("id")
	id, err := strconv.Atoi(bandId)
	helper.ErrorPanic(err)
	updateBandRequest.Id = id
	controller.bandService.Update(updateBandRequest)
	w.Header().Add("Content-Type", "application/json")
}

// @Summary returns a list of all bands
// @Description returns a list of all bands
// @Tags Bands
// @Router /bands [get]
func (controller *BandController) FindAll(w http.ResponseWriter, r *http.Request) {
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
	bandId := r.PathValue("id")
	id, err := strconv.Atoi(bandId)
	helper.ErrorPanic(err)

	bandResponse := controller.bandService.FindById(id)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bandResponse)
}

// @Summary deletes a band by its id
// @Description deletes a band by its id
// @Tags Bands
// @Param bandId path string true "find band by id"
// @Router /bands/{bandId} [delete]
func (controller *BandController) Delete(w http.ResponseWriter, r *http.Request) {
	bandId := r.PathValue("id")
	id, err := strconv.Atoi(bandId)
	helper.ErrorPanic(err)

	controller.bandService.Delete(id)

	w.Header().Add("Content-Type", "application/json")
}
