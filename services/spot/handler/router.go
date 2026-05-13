package handler

import (
	"smart-parking/services/spot/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	logger "github.com/sirupsen/logrus"
)

func NewRouter(zoneSvc *service.ZoneService, logger *logger.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	zoneHandler := NewZoneHandler(zoneSvc, logger)

	router.Post("/zones", zoneHandler.Create)
	return router
}
