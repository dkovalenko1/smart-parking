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

func (s *ZoneService) CreateParkingZone(ctx context.Context, args CreateParkingZoneArgs) (*model.Zone, error) {
	s.logger.Info("creating parking zone ", args.Name)
	zone, err := model.CreatNewZone(args.Name, args.Description)
	if err != nil {
		return nil, fmt.Errorf("error creating parking zone: %w", err)
	}
	s.logger.Info("parking zone ", args.Name, "created")

	s.logger.Info("saving parking zone ", args.Name)
	if err := s.zones.Save(ctx, zone); err != nil {
		return nil, fmt.Errorf("error saving parking zone : %w", err)
	}
	s.logger.Info("parking zone ", args.Name, " saved")

	return zone, nil
}

func (s *ZoneService) GetParkingZones(ctx context.Context) ([]model.Zone, error) {
	s.logger.Info("getting parking zones")
	zones, err := s.zones.GetZones(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting parking zones: %w", err)
	}

	return zones, nil
}
