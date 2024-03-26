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

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(loadBands())
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
