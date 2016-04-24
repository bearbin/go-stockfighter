package sflib

import (
	"fmt"
	"encoding/json"
)

type VenueHeartbeatResponse struct {
	APIStatusResponse
	Venue string `json:"venue"`
}

// VenueHeartbeat checks if a venue is up and running.
func (c *Client) VenueHeartbeat(venue string) (*VenueHeartbeatResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("venues/%s/heartbeat", venue)
	data, err := c.call("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &VenueHeartbeatResponse{}
	err = json.NewDecoder(*data).Decode(response)
	if err != nil {
		return nil, err
	}

	// Check if the request completed successfully.
	err = response.CheckAPIStatus()
	if err != nil {
		return nil, err
	}

	// If nothing failed, then return the response.
	return response, nil
}

type VenueSymbol struct {
	Name string `json:"name"`
	Symbol string `json:"symbol"`
}

type VenueStocksResponse struct {
	APIStatusResponse
	Symbols []VenueSymbol `json:"symbols"`
}

// VenueHeartbeat checks if a venue is up and running.
func (c *Client) VenueStocks(venue string) (*VenueStocksResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("venues/%s/stocks", venue)
	data, err := c.call("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &VenueStocksResponse{}
	err = json.NewDecoder(*data).Decode(response)
	if err != nil {
		return nil, err
	}

	// Check if the request completed successfully.
	err = response.CheckAPIStatus()
	if err != nil {
		return nil, err
	}

	// If nothing failed, then return the response.
	return response, nil
}
