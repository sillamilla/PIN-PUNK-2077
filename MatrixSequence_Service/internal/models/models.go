package models

const (
	MatrixWidth  = 5
	MatrixHeight = 5

	KeySequenceLength = 7
)

type MatrixData struct {
	Matrix
	KeySequence
}

type Matrix struct {
	Data [MatrixHeight][MatrixWidth]int `json:"matrix"` //todo check
}

type KeySequence struct {
	Keys []int `json:"keys"`
}
