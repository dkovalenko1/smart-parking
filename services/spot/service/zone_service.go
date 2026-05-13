package service

import (
	"context"
	"fmt"
	"smart-parking/services/spot/model"

	logger "github.com/sirupsen/logrus"
)

type ZoneService struct {
	zones  ZoneRepository
	logger *logger.Logger
}

func NewZoneService(zones ZoneRepository, logger *logger.Logger) *ZoneService {
	return &ZoneService{
		logger: logger,
		zones:  zones,
	}
}

func (p *ZoneService) CreateParkingZone(ctx context.Context, args CreateParkingZoneArgs) (*model.Zone, error) {
	p.logger.Debug("creating parking zone ", args.Name)
	zone, err := model.CreatNewZone(args.Name, args.Description)
	if err != nil {
		return nil, fmt.Errorf("error creating parking zone: %w", err)
	}
	p.logger.Debug("parking zone ", args.Name, "created")

	p.logger.Debug("saving parking zone", args.Name)
	if err := p.zones.Save(ctx, zone); err != nil {
		return nil, fmt.Errorf("error saving parking zone : %w", err)
	}
	p.logger.Debug("parking zone ", args.Name, "saved")

	return zone, nil
}
