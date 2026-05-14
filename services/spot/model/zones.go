package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Zone struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

func CreatNewZone(name, description string) (zone *Zone, err error) {
	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if name == "" {
		return nil, ErrEmptyZoneName
	}
	if description == "" {
		return nil, ErrEmptyZoneDescription
	}

	now := time.Now().UTC()
	return &Zone{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedAt:   now,
	}, err
}
