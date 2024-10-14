package helper

import (
	"MiniGame-PinUp/MatrixSequence_Service/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(status)
	io.WriteString(w, errMsg)
}

func PostJSON(url string, data models.MatrixData) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error serializing data: %w", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	return resp, nil
}

func ReadJSONResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("error decoding JSON response: %w", err)
	}

	return nil
}
