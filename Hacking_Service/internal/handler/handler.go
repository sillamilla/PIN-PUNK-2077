package handler

import (
	"MiniGame-PinUp/Hacking_Service/internal/models"
	"MiniGame-PinUp/Hacking_Service/internal/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Handler struct {
	service service.MatrixService
}

func New(service service.MatrixService) Handler {
	return Handler{service: service}
}

func (h *Handler) Hack(w http.ResponseWriter, r *http.Request) {
	readAll, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Error reading request body: %v", err))
		return
	}
	defer r.Body.Close()

	var data models.MatrixData
	if err = json.Unmarshal(readAll, &data); err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Error unmarshalling JSON: %v", err))
		return
	}

	status, err := h.service.Hack(r.Context(), data.Matrix, data.KeySequence)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Error hacking matrix: %v", err))
		return
	}

	response := models.StatusResponse{Status: status}
	sendResponse(w, response, http.StatusOK)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	val, err := h.service.GetAll(r.Context())
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	if err = json.NewEncoder(w).Encode(val); err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Error encoding JSON: %v", err))
		return
	}
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
