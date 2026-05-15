package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type SpotType string

const (
	SpotTypeRegular  SpotType = "regular"
	SpotTypeEV       SpotType = "ev"
	SpotTypeDisabled SpotType = "disabled"
)

type SpotStatus string

const (
	SpotStatusAvailable   SpotStatus = "available"
	SpotStatusReserved    SpotStatus = "reserved"
	SpotStatusOccupied    SpotStatus = "occupied"
	SpotStatusMaintenance SpotStatus = "maintenance"
)

func ParseSpotType(spotType string) (SpotType, error) {
	value := strings.TrimSpace(strings.ToLower(spotType))
	switch SpotType(value) {
	case SpotTypeRegular, SpotTypeEV, SpotTypeDisabled:
		return SpotType(value), nil
	default:
		return "", ErrInvalidSpotType
	}
}

type Spot struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	ZoneID    uuid.UUID  `json:"zoneId" db:"zone_id"`
	Number    string     `json:"number" db:"number"`
	Type      SpotType   `json:"type" db:"type"`
	Status    SpotStatus `json:"status" db:"status"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
}

func CreateNewSpot(zoneId uuid.UUID, number, spotTypeString string) (spot *Spot, err error) {
	number, spotTypeString = strings.TrimSpace(number), strings.TrimSpace(spotTypeString)
	if number == "" {
		return nil, ErrEmptySpotNumber
	}
	if spotTypeString == "" {
		return nil, ErrEmptySpotType
	}
	spotType, err := ParseSpotType(spotTypeString)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	return &Spot{
		ID:        uuid.New(),
		ZoneID:    zoneId,
		Number:    number,
		Type:      spotType,
		Status:    SpotStatusAvailable,
		CreatedAt: now,
	}, err
}
