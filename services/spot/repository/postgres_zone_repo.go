package repository

import (
	"context"
	"errors"
	"smart-parking/services/spot/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	logger "github.com/sirupsen/logrus"
)

type ZoneRepository struct {
	pool   *pgxpool.Pool
	logger *logger.Logger
}

func NewZoneRepository(pool *pgxpool.Pool, logger *logger.Logger) *ZoneRepository {
	return &ZoneRepository{pool: pool, logger: logger}
}

func (repo *ZoneRepository) Save(ctx context.Context, zone *model.Zone) error {
	_, err := repo.pool.Exec(ctx, `
		INSERT INTO parking_zones (id, name, description, created_at) 
		VALUES ($1, $2, $3, $4)
`,
		zone.ID,
		zone.Name,
		zone.Description,
		zone.CreatedAt,
	)
	if err != nil {
		repo.logger.Error("error on saving parking zones: ", err)
	}
	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
		if pgErr.Code == "23505" {
			return model.ErrZoneAlreadyExists
		}
	}
	return err
}

func (repo *ZoneRepository) GetZones(ctx context.Context) ([]model.Zone, error) {
	rows, err := repo.pool.Query(ctx, `SELECT id, name, description, created_at FROM parking_zones`)
	if err != nil {
		repo.logger.Error("error on listing parking zones: ", err)
		return nil, err
	}
	defer rows.Close()

	zones, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Zone])
	if err != nil {
		repo.logger.Error("error on mapping rows to zone struct: ", err)
		return nil, err
	}

	return zones, nil
}
