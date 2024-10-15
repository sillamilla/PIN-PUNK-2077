package handler

import (
	"MiniGame-PinUp/MatrixSequence_Service/internal/service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Handler struct {
	service service.MatrixService
}

func New(service service.MatrixService) Handler {
	return Handler{service: service}
}

func (h *Handler) GenerateMatrixData(w http.ResponseWriter, r *http.Request) {
	data := h.service.NewMatrixData(r.Context())
	sendResponse(w, data, http.StatusCreated)
}

func (h *Handler) CallHack(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.HackMatrix(r.Context())
	if err != nil {
		if errors.Is(err, errors.New("matrix data is empty")) {
			sendError(w, http.StatusBadRequest, fmt.Sprintf("Error hacking matrix: %v", err))
			return
		}
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Error hacking matrix: %v", err))
		return
	}

	sendResponse(w, status, http.StatusOK)
}

func sendError(w http.ResponseWriter, status int, errMsg string) {
	if status == http.StatusInternalServerError {
		errMsg = "Internal server error"
	}

	w.WriteHeader(status)
	w.Write([]byte(errMsg))
}

func sendResponse(w http.ResponseWriter, data interface{}, status int) {
	dataMarshal, err := json.Marshal(data)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Error marshalling JSON: %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(dataMarshal)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Error writing response: %v", err))
		return
	}
}
