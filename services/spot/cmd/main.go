package main

import (
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.SetReportCaller(true)

	dbURL := mustEnv("DATABASE_URL")

	if err := runMigrations(dbURL); err != nil {
		logger.Error("migrations failed: ", err)
		os.Exit(1)
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
		logger.Error("missing required env variable: ", "key", key)
		os.Exit(1)
	}
	return v
}
