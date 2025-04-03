package client

// ModelResponse is the response for a model
type ModelResponse struct {
	// ID is the ID of the model
	ID string `json:"id"`

	// Tags are the tags of the model
	Tags []string `json:"tags"`

	// Files are the files of the model
	Files []string `json:"files"`

	// Created is the creation time of the model
	Created int64 `json:"created"`
}

// Client is the client for the Docker Model Runner
type Client struct {
	BaseURL string
}

// NewClient creates a new client for the Docker Model Runner
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
	}
}
