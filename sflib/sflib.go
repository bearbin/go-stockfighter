package sflib

import (
	"http"
	"net/url"
)

const (
	libraryVersion = "0.1"
	userAgent      = "bearbin-sflib/" + libraryVersion

	// Default Base URL for the API
	baseURL = "https://api.stockfighter.io/ob/api/"
)

// A Client manages the connection with the stockfighter API.
type Client struct {
	// Well, we need a HTTP client at least.
	client *http.Client

	// Base URL for API requests, which should be provided with a trailing slash.
	BaseURL *url.URL

	// User agent for requests to the API.
	UserAgent string

	// API token to authenticate with the API.
	APIToken string
}

// NewClient provides a new client with default values.
func NewClient(apiToken string) *Client {
	parsedBaseURL, _ = url.Parse(baseURL)

	return &Client{client: http.DefaultClient, BaseURL: parsedBaseURL, userAgent: userAgent, APIToken: apiToken}
}

// Call simply sends a HTTP request
func (c *Client) Call(method string, endpoint string, data string) {
	// Create the HTTP request.
	req, err := http.NewRequest(method, c.BaseURL+endpoint, nil)

	// Add authorisation header with API token.
	req.Header.Add("X-Starfighter-Authorization", c.APIToken)

	// Do the request.
	response, err := c.client.Do(req)
}
