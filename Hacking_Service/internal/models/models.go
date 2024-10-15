package models

import "time"

const (
	MatrixWidth  = 5
	MatrixHeight = 5

	KeySequenceLength = 7
)

type StatusResponse struct {
	Status string `json:"status"`
}

type MatrixData struct {
	Matrix
	KeySequence
}

type Matrix struct {
	Data [MatrixHeight][MatrixWidth]int `json:"matrix"`
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
