package handler

import (
	"smart-parking/services/spot/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	logger "github.com/sirupsen/logrus"
)

func NewRouter(zoneSvc *service.ZoneService, spotSvc *service.SpotService, pool *pgxpool.Pool, logger *logger.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	zoneHandler := NewZoneHandler(zoneSvc, logger)
	spotHandler := NewSpotHandler(spotSvc, logger)
	healthHandler := NewHealthHandler(pool, logger)

	router.Post("/zones", zoneHandler.Create)
	router.Get("/health", healthHandler.Check)
	router.Get("/zones", zoneHandler.Get)
	router.Get("/zones/{id}", zoneHandler.GetById)
	router.Post("/spots", spotHandler.Create)

	return router
}
