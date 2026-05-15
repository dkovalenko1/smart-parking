package repository

import (
	"context"
	"errors"
	"smart-parking/services/spot/model"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	logger "github.com/sirupsen/logrus"
)

type SpotRepository struct {
	pool   *pgxpool.Pool
	logger *logger.Logger
}

func NewSpotRepository(pool *pgxpool.Pool, logger *logger.Logger) *SpotRepository {
	return &SpotRepository{pool: pool, logger: logger}
}

func (repo *SpotRepository) Save(ctx context.Context, spot *model.Spot) error {
	_, err := repo.pool.Exec(ctx, `
		INSERT INTO spots (id, zone_id, number, type, status, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
`,
		spot.ID,
		spot.ZoneID,
		spot.Number,
		spot.Type,
		spot.Status,
		spot.CreatedAt,
	)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				return model.ErrSpotAlreadyExists
			}
			if pgErr.Code == "23503" {
				return model.ErrZoneNotFound
			}
		}
		repo.logger.Error("error on saving spot: ", err)
		return err
	}

	return nil
}
