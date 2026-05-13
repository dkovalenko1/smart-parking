package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"smart-parking/services/spot/handler"
	"smart-parking/services/spot/repository"
	"smart-parking/services/spot/service"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

func main() {
	var logger = log.New()
	logger.SetReportCaller(true)

	dbURL := mustEnv("DATABASE_URL")
	port := mustEnv("PORT")

	if err := runMigrations(dbURL); err != nil {
		logger.Fatal("migrations failed: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		logger.Fatal(err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		logger.Fatal("database ping failed: ", err)
	}
	logger.Info("database ping succeeded")

	zoneRepo := repository.NewZoneRepository(pool)
	zoneService := service.NewZoneService(zoneRepo, logger)
	router := handler.NewRouter(zoneService, logger)

	logger.Info("starting server")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		logger.Fatal("server error: ", err)
	}
}

func runMigrations(dbURL string) error {
	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		return err
	}
	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatal("missing required env variable: ", key)
	}
	return v
}
