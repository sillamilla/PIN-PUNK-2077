package hackService

type MatrixData struct {
	Matrix
	KeySequence
}

const (
	MatrixWidth  = 5
	MatrixHeight = 5
)

type Matrix struct {
	Data [MatrixHeight][MatrixWidth]int `json:"matrix"`
}

type KeySequence struct {
	Keys []int `json:"keys"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
