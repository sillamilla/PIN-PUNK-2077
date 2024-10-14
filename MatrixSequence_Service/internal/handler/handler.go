package handler

import (
	"MiniGame-PinUp/MatrixSequence_Service/internal/helper"
	"MiniGame-PinUp/MatrixSequence_Service/internal/models"
	"MiniGame-PinUp/MatrixSequence_Service/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

var matrixData models.MatrixData // Глобальная переменная для хранения matrixData

type handler struct {
	service service.Matrix
}

func NewHandler(service service.Matrix) handler {
	return handler{service: service}
}

func (h *handler) GetSequence(w http.ResponseWriter, r *http.Request) {
	matrixData = h.service.NewMetricData()

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(matrixData); err != nil {
		helper.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Error encoding JSON: %v", err))
		return
	}
}

func (h *handler) CallHack(w http.ResponseWriter, r *http.Request) {
	if len(matrixData.Data) == 0 || len(matrixData.Keys) == 0 {
		helper.SendError(w, http.StatusBadRequest, fmt.Sprintf("No metric or sequence found, please generate them first."))
		return
	}

	jsonResp, err := helper.PostJSON("http://localhost:8081/hack", matrixData)
	if err != nil {
		helper.SendError(w, http.StatusBadRequest, fmt.Sprintf("Error sending request: %w", err))
		return
	}

	var result models.StatusResponse
	err = helper.ReadJSONResponse(jsonResp, &result)
	if err != nil {
		helper.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Error reading response: %w", err))
		return
	}

	if err = json.NewEncoder(w).Encode(result); err != nil {
		helper.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Error encoding JSON: %v", err))
		return
	}
}
