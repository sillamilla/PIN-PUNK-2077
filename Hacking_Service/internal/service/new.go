package service

import (
	"MiniGame-PinUp/Hacking_Service/internal/repository"
)

type Service struct {
	Matrix
}

func New(repository *repository.Repository) *Service {
	return &Service{NewService(repository)}
}
