package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListModels lists all models
func (c *Client) ListModels() ([]ModelResponse, error) {
	var models []ModelResponse

	httpClient := http.Client{}
	resp, err := httpClient.Get(c.BaseURL + "/models")
	if err != nil {
		return models, fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models, fmt.Errorf("read all: %w", err)
	}

	err = json.Unmarshal(body, &models)
	if err != nil {
		return models, fmt.Errorf("json unmarshal: %w", err)
	}

	return models, nil
}
