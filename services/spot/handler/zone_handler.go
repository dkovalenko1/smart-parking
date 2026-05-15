package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"smart-parking/services/spot/model"
	"smart-parking/services/spot/service"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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
	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		writeError(writer, h.logger, http.StatusBadRequest, "invalid request body")
		return
	}

	zone, err := h.service.CreateParkingZone(req.Context(), service.CreateParkingZoneArgs{
		Name:        request.Name,
		Description: request.Description,
	})
	if err != nil {
		switch {
		case errors.Is(err, model.ErrEmptyZoneName):
			writeError(writer, h.logger, http.StatusBadRequest, "invalid request body, name is required")
		case errors.Is(err, model.ErrEmptyZoneDescription):
			writeError(writer, h.logger, http.StatusBadRequest, "invalid request body, zone description is required")
		case errors.Is(err, model.ErrZoneAlreadyExists):
			writeError(writer, h.logger, http.StatusConflict, "parking zone already exists")
		default:
			h.logger.Error("failed to create parking zone: ", err)
			writeError(writer, h.logger, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	writeJSON(writer, h.logger, http.StatusCreated, zone)
}

func (h *ZoneHandler) Get(writer http.ResponseWriter, req *http.Request) {
	zones, err := h.service.GetParkingZones(req.Context())
	if err != nil {
		h.logger.Error("failed to get parking zones: ", err)
		writeError(writer, h.logger, http.StatusInternalServerError, "internal server error")
		return
	}

	writeJSON(writer, h.logger, http.StatusOK, zones)
}

func (h *ZoneHandler) GetById(writer http.ResponseWriter, req *http.Request) {
	idString := chi.URLParam(req, "id")
	id, err := uuid.Parse(idString)
	if err != nil {
		writeError(writer, h.logger, http.StatusBadRequest, "invalid parking zone id")
		return
	}

	zone, err := h.service.GetParkingZoneById(req.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrZoneNotFound):
			writeError(writer, h.logger, http.StatusNotFound, "parking zone not found")
		default:
			h.logger.Error("failed to get parking zone by id: ", err)
			writeError(writer, h.logger, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	writeJSON(writer, h.logger, http.StatusOK, zone)
}
