package models

import "time"

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

type HackData struct {
	Key  string   `pg:"key,unique"`
	Data JSONData `pg:"data,type:jsonb"`
}

type JSONData struct {
	Created      time.Time `json:"created"`
	ResultOfHack int       `json:"result_of_hack"`
}
