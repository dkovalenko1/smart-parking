package handler

import (
	"github.com/google/uuid"
)

type CreateZoneRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateSpotRequest struct {
	ZoneID uuid.UUID `json:"zoneId"`
	Number string    `json:"number"`
	Type   string    `json:"type"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
