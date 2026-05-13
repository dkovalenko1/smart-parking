package repository

import (
	"context"
	"smart-parking/services/spot/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ZoneRepository struct {
	pool *pgxpool.Pool
}

func NewZoneRepository(pool *pgxpool.Pool) *ZoneRepository {
	return &ZoneRepository{pool: pool}
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
	return err
}
