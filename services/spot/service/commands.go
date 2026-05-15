package service

import "github.com/google/uuid"

type CreateParkingZoneArgs struct {
	Name        string
	Description string
}

type CreateParkingSpotArgs struct {
	ZoneID uuid.UUID
	Number string
	Type   string
}
