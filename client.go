package tumblr2go

import (
	"github.com/anpez/tumblr2go/interfaces"
)

const API_BASE_URL = "https://api.tumblr.com/v2"

type Client struct {
	httpClient interfaces.HttpClient
	apiKey     string
}

// Creates a new client for Tumblr API v2 using the provided api key.
func NewClient(apiKey string) *Client {
	return &Client{
		httpClient: newHttpClient(),
		apiKey:     apiKey,
	}
}

// Creates a new client for Tumblr API v2 using the provided api key
// and the provided HttpClient (good for testing).
func NewClientWithHttp(apiKey string, httpClient interfaces.HttpClient) *Client {
	return &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}
