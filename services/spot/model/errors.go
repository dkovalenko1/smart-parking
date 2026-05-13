package model

import "errors"

var (
	ErrEmptyZoneName        = errors.New("zone name is empty")
	ErrEmptyZoneDescription = errors.New("zone description is empty")
	ErrZoneAlreadyExists    = errors.New("zone already exists")
)
