package handler

import (
	"encoding/json"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func writeError(writer http.ResponseWriter, logger *logger.Logger, status int, msg string) {
	writeJSON(writer, logger, status, ErrorResponse{Error: msg})
}

func writeJSON(writer http.ResponseWriter, logger *logger.Logger, status int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(data); err != nil {
		logger.Error("failed to write response: ", err)
	}
}
