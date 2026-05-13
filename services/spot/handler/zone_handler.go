package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"smart-parking/services/spot/model"
	"smart-parking/services/spot/service"

	logger "github.com/sirupsen/logrus"
)

type ZoneHandler struct {
	service *service.ZoneService
	logger  *logger.Logger
}

func NewZoneHandler(service *service.ZoneService, logger *logger.Logger) *ZoneHandler {
	return &ZoneHandler{service: service, logger: logger}
}

func (h *ZoneHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var request CreateZoneRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		h.writeError(writer, http.StatusBadRequest, "invalid request body")
		return
	}

	zone, err := h.service.CreateParkingZone(req.Context(), service.CreateParkingZoneArgs{
		Name:        request.Name,
		Description: request.Description,
	})
	if err != nil {
		switch {
		case errors.Is(err, model.ErrEmptyZoneName):
			h.writeError(writer, http.StatusBadRequest, "invalid request body, name is required")
		case errors.Is(err, model.ErrEmptyZoneDescription):
			h.writeError(writer, http.StatusBadRequest, "invalid request body, zone description is required")
		case errors.Is(err, model.ErrZoneAlreadyExists):
			h.writeError(writer, http.StatusConflict, "parking zone already exists")
		default:
			h.logger.Error("failed to create parking zone: ", err)
			h.writeError(writer, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	h.writeJSON(writer, http.StatusCreated, toZoneResponse(zone))
}

func toZoneResponse(zone *model.Zone) CreateZoneResponse {
	return CreateZoneResponse{
		ID:          zone.ID,
		Name:        zone.Name,
		Description: zone.Description,
		CreatedAt:   zone.CreatedAt,
	}
}

func (h *ZoneHandler) writeError(writer http.ResponseWriter, status int, msg string) {
	h.writeJSON(writer, status, ErrorResponse{Error: msg})
}

func (h *ZoneHandler) writeJSON(writer http.ResponseWriter, status int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		h.logger.Error("failed to write response: ", err)
	}
}
