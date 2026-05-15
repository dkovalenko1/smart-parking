package service

import (
	"context"
	"fmt"
	"smart-parking/services/spot/model"

	logger "github.com/sirupsen/logrus"
)

type SpotService struct {
	spots  SpotRepository
	logger *logger.Logger
}

func NewSpotService(spots SpotRepository, logger *logger.Logger) *SpotService {
	return &SpotService{
		logger: logger,
		spots:  spots,
	}
}

func (s *SpotService) CreateParkingSpot(ctx context.Context, args CreateParkingSpotArgs) (*model.Spot, error) {
	s.logger.Info("creating parking spot")
	spot, err := model.CreateNewSpot(args.ZoneID, args.Number, args.Type)
	if err != nil {
		return nil, fmt.Errorf("error creating parking spot: %w", err)
	}
	s.logger.Info("parking spot ", args.Number, " created")

	s.logger.Info("saving parking spot ", args.Number)
	if err := s.spots.Save(ctx, spot); err != nil {
		return nil, fmt.Errorf("error saving parking spot : %w", err)
	}
	s.logger.Info("parking spot ", args.Number, " saved")

	return spot, nil
}
