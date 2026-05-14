package service

import (
	"context"
	"smart-parking/services/spot/model"
)

type ZoneRepository interface {
	Save(ctx context.Context, zone *model.Zone) error
	GetZones(ctx context.Context) ([]model.Zone, error)
}
