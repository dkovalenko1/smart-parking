package model

import "errors"

var (
	ErrEmptyZoneName        = errors.New("zone name is empty")
	ErrEmptyZoneDescription = errors.New("zone description is empty")
	ErrZoneAlreadyExists    = errors.New("zone already exists")
	ErrZoneNotFound         = errors.New("zone not found")
	ErrEmptySpotNumber      = errors.New("spot number is empty")
	ErrEmptySpotType        = errors.New("spot type is empty")
	ErrInvalidSpotType      = errors.New("spot type is invalid")
	ErrSpotAlreadyExists    = errors.New("spot already exists")
)
