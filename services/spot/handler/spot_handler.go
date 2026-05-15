package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"smart-parking/services/spot/model"
	"smart-parking/services/spot/service"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type SpotHandler struct {
	service *service.SpotService
	logger  *logger.Logger
}

func NewSpotHandler(service *service.SpotService, logger *logger.Logger) *SpotHandler {
	return &SpotHandler{service: service, logger: logger}
}

func (h *SpotHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var request CreateSpotRequest

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		writeError(writer, h.logger, http.StatusBadRequest, "invalid request body")
		return
	}

	if request.ZoneID == uuid.Nil { //if zoneId is not specified, since our struct uses uuid.UUID, it will be passed through as uuid.Nil (00000000-00...)
		writeError(writer, h.logger, http.StatusBadRequest, "invalid request body, zoneId is required")
		return
	}

	spot, err := h.service.CreateParkingSpot(req.Context(), service.CreateParkingSpotArgs{
		ZoneID: request.ZoneID,
		Number: request.Number,
		Type:   request.Type,
	})
	if err != nil {
		switch {
		case errors.Is(err, model.ErrZoneNotFound):
			writeError(writer, h.logger, http.StatusNotFound, "specified parking zone not found")
		case errors.Is(err, model.ErrEmptySpotNumber):
			writeError(writer, h.logger, http.StatusBadRequest, "invalid request body, number is required")
		case errors.Is(err, model.ErrEmptySpotType):
			writeError(writer, h.logger, http.StatusBadRequest, "invalid request body, type is required")
		case errors.Is(err, model.ErrInvalidSpotType):
			writeError(writer, h.logger, http.StatusBadRequest, "invalid request body, type is invalid")
		case errors.Is(err, model.ErrSpotAlreadyExists):
			writeError(writer, h.logger, http.StatusConflict, "parking spot already exists")
		default:
			h.logger.Error("failed to create parking spot: ", err)
			writeError(writer, h.logger, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	writeJSON(writer, h.logger, http.StatusCreated, spot)
}
