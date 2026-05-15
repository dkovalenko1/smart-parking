package service

import (
	"context"
	"smart-parking/services/spot/model"

	"github.com/google/uuid"
)

type ZoneRepository interface {
	Save(ctx context.Context, zone *model.Zone) error
	GetZones(ctx context.Context) ([]model.Zone, error)
	GetZoneById(ctx context.Context, id uuid.UUID) (*model.Zone, error)
}

type SpotRepository interface {
	Save(ctx context.Context, zone *model.Spot) error
}
