package models

type JsonResponse struct {
	Error   bool   `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type MatrixData struct {
	Matrix
	KeySequence
}

type Matrix struct {
	Data [5][5]int `json:"matrix"`
}

type KeySequence struct {
	Keys []int `json:"keys"`
}
