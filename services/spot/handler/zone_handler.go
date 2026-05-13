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
		writeError(writer, http.StatusBadRequest, "invalid request body")
		return
	}

	zone, err := h.service.CreateParkingZone(req.Context(), service.CreateParkingZoneArgs{
		Name:        request.Name,
		Description: request.Description,
	})
	if err != nil {
		switch {
		case errors.Is(err, model.ErrEmptyZoneName):
			writeError(writer, http.StatusBadRequest, "invalid request body, name is required")
		default:
			h.logger.Error("failed to create parking zone: ", err)
			writeError(writer, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	writeJSON(writer, http.StatusCreated, toZoneResponse(zone))
}

func toZoneResponse(zone *model.Zone) CreateZoneResponse {
	return CreateZoneResponse{
		ID:          zone.ID,
		Name:        zone.Name,
		Description: zone.Description,
		CreatedAt:   zone.CreatedAt,
	}
}

func writeError(writer http.ResponseWriter, status int, msg string) {
	writeJSON(writer, status, ErrorResponse{Error: msg})
}

func writeJSON(writer http.ResponseWriter, status int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		logger.Error("failed to write response: ", err)
	}
}
