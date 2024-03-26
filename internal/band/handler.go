package band

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct{}

func NewHandlerFunc() *Handler {
	return &Handler{}
}

// @Summary returns a list of all bands
// @Description returns a list of all bands
// @Tags Bands
// @Router /band [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getBands())
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	log.Println("handling READ request - Method:", r.Method)
	band, exists := loadBands()[r.PathValue("id")]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(band)
}
