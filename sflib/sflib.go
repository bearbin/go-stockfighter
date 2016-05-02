package sflib

import (
	"io"
	"net/http"
)

const (
	libraryVersion   = "0.1"
	defaultUserAgent = "bearbin-sflib/" + libraryVersion

	// Default Base URL for the API
	baseURL = "https://api.stockfighter.io/"
)

// A Client manages the connection with the stockfighter API.
// It is recommended to use NewClient to get a new Client object with sensible
// default values.
type Client struct {
	// Well, we need a HTTP client at least.
	client *http.Client

	// Base URL for API requests, which should be provided with a trailing slash.
	BaseURL string

	// User agent for requests to the API.
	UserAgent string

	// API token to authenticate with the API.
	APIToken string
}

// NewClient provides a new client with default values.
func NewClient(apiToken string) *Client {
	return &Client{client: http.DefaultClient, BaseURL: baseURL, UserAgent: defaultUserAgent, APIToken: apiToken}
}

// call sends a request to the APi.
// It uses the specified method, endpoint and data, and uses the Base URL, User
// Agent and API token from the Client object it is associated with.
func (c *Client) call(method string, endpoint string, data io.Reader) (*io.ReadCloser, error) {
	// Create the HTTP request.
	requestPath := c.BaseURL + endpoint
	req, err := http.NewRequest(method, requestPath, nil)
	if data != nil {
		req, err = http.NewRequest(method, requestPath, data)
	}

	if err != nil {
		return nil, err
	}

	// Set the user agent.
	req.Header.Set("User-Agent", c.UserAgent)

	// Add authorisation header with API token.
	req.Header.Add("X-Starfighter-Authorization", c.APIToken)

	// Do the request.
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the information, assume that the API did not error out.
	return &response.Body, nil
}
