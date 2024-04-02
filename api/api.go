package api

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

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	db := db.InitializeDatabase()
	validate := validator.New()
	bandRepository := repository.NewBandRepositoryImpl(db)
	bandService := service.NewBandServiceImpl(bandRepository, validate)
	bandController := controller.NewBandController(bandService)
	routes := router.NewRouter(bandController)

	server := http.Server{
		Addr:    s.addr,
		Handler: RequestLoggerMiddleware(routes),
	}

	log.Info().Msgf("Server started, listening on port %s", s.addr)

	return server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msgf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
