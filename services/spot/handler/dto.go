package handler

import (
	"time"

	"github.com/google/uuid"
)

type CreateZoneRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateZoneResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetZoneByIdRequest struct {
	ID uuid.UUID `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
