package handler

import (
	"MiniGame-PinUp/Hacking_Service/internal/helper"
	"MiniGame-PinUp/Hacking_Service/internal/models"
	"MiniGame-PinUp/Hacking_Service/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type handler struct {
	service service.Matrix
}

func NewHandler(service service.Matrix) handler {
	return handler{service: service}
}

func (h *handler) Hack(w http.ResponseWriter, r *http.Request) {
	var data models.MatrixData

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		helper.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Error decoding JSON from request body: %v", err))
		return
	}

	keys, number := h.service.CountMatches(data.Matrix, data.KeySequence)
	if len(keys) != 0 {
		//save in success case
		err := h.service.SaveHackAttempt(keys, number)
		if err != nil {
			helper.SendError(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
			return
		}
	}

	status := "fail"
	if number > 0 {
		status = "success"
	}
	response := models.StatusResponse{Status: status}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		helper.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Error encoding JSON: %v", err))
		return
	}
}

func (h handler) GetAll(w http.ResponseWriter, r *http.Request) {
	val, err := h.service.GetAll()
	if err != nil {
		helper.SendError(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	if err = json.NewEncoder(w).Encode(val); err != nil {
		helper.SendError(w, http.StatusInternalServerError, fmt.Sprintf("Error encoding JSON: %v", err))
		return
	}
}
