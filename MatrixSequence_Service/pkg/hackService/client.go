package hackService

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client interface {
	HackMatrix(ctx context.Context, matrix MatrixData) (*StatusResponse, error)
}

type client struct {
	httpClient   http.Client
	addr         string
	hackEndpoint string
}

func NewClient(addr, hackEndpoint string) Client {
	return &client{
		httpClient: http.Client{
			Timeout: time.Second * 10,
		},
		addr:         addr,
		hackEndpoint: hackEndpoint,
	}
}

func (c *client) HackMatrix(ctx context.Context, matrix MatrixData) (*StatusResponse, error) {
	req, err := json.Marshal(matrix)
	if err != nil {
		return nil, fmt.Errorf("error marshalling matrix: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.addr+c.hackEndpoint, bytes.NewReader(req))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	readAll, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s status code: %d", string(readAll), res.StatusCode)
	}

	var response *StatusResponse
	if err = json.Unmarshal(readAll, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return response, nil
}
