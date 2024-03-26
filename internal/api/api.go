package api

import (
	"band_api/internal/band"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	handler := &band.Handler{}

	router.HandleFunc("GET /band", handler.GetAll)
	router.HandleFunc("GET /band/{id}", handler.FindByID)

	return http.ListenAndServe(s.addr, router)
}
