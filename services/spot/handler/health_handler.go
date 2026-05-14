package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	logger "github.com/sirupsen/logrus"
)

type HealthHandler struct {
	pool   *pgxpool.Pool
	logger *logger.Logger
}

func NewHealthHandler(pool *pgxpool.Pool, logger *logger.Logger) *HealthHandler {
	return &HealthHandler{pool: pool, logger: logger}
}

func (h *HealthHandler) Check(writer http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.pool.Ping(ctx)
	if err != nil {
		writeError(writer, h.logger, http.StatusServiceUnavailable, "database unavailable")
		h.logger.Error("health check failed: ", err)
		return
	}
	writeJSON(writer, h.logger, http.StatusOK, map[string]string{"status": "ok"})
}
