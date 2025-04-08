package client

import (
	"fmt"
	"io"
	"net/http"
)

// DeleteModel deletes a model by namespace and name
func (c *Client) DeleteModel(namespace string, name string) ([]byte, error) {
	var bytes []byte

	reqURL := c.BaseURL + "/models/" + namespace + "/" + name

	httpClient := http.Client{}
	req, err := http.NewRequest("DELETE", reqURL, nil)
	if err != nil {
		return bytes, fmt.Errorf("create delete request (%s): %w", reqURL, err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return bytes, fmt.Errorf("do request (%s): %w", reqURL, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return bytes, fmt.Errorf("read all: %w", err)
	}
	bytes = body

	return bytes, nil
}
