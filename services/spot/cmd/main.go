package main

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.SetReportCaller(true)

	dbURL := mustEnv("DATABASE_URL")

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
		logger.Fatal("missing required env variable: ", key)
	}
	return v
}
