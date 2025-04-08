package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// CreateModel creates a model
func (c *Client) CreateModel(fqmn string) ([]byte, error) {
	var bytes []byte

	payload := fmt.Sprintf(`{"from": "%s"}`, fqmn)

	httpClient := http.Client{}
	resp, err := httpClient.Post(c.BaseURL+"/models/create", "application/json", strings.NewReader(payload))
	if err != nil {
		return bytes, fmt.Errorf("http post: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return bytes, fmt.Errorf("read all: %w", err)
	}
	bytes = body
	return bytes, nil
}
