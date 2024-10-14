package service

type Service struct {
	Matrix
}

func New() *Service {
	return &Service{NewService()}
}
