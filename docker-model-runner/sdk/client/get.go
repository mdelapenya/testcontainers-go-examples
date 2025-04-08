package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetModel gets a model by namespace and name
func (c *Client) GetModel(namespace string, name string) (ModelResponse, error) {
	var model ModelResponse

	httpClient := http.Client{}
	resp, err := httpClient.Get(c.BaseURL + "/models/" + namespace + "/" + name)
	if err != nil {
		return model, fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model, fmt.Errorf("read all: %w", err)
	}

	err = json.Unmarshal(body, &model)
	if err != nil {
		return model, fmt.Errorf("json unmarshal: %w", err)
	}

	return model, nil
}
