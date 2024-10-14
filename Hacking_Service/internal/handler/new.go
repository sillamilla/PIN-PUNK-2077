package handler

import (
	"MiniGame-PinUp/Hacking_Service/internal/service"
)

type Handler struct {
	handler
}

func New(service *service.Service) *Handler {
	return &Handler{NewHandler(service)}
}
