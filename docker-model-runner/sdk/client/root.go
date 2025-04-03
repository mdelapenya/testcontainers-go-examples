package client

import (
	"fmt"
	"io"
	"net/http"
)

// Root gets the root of the Docker Model Runner
func (c *Client) Root() ([]byte, error) {
	var bytes []byte

	httpClient := http.Client{}
	resp, err := httpClient.Get(c.BaseURL)
	if err != nil {
		return bytes, fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return bytes, fmt.Errorf("read all: %w", err)
	}
	bytes = body

	return bytes, nil
}
