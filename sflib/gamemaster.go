package sflib

import (
	"encoding/json"
	"fmt"
)

// LevelStartResponse is a response from the gamemaster API.
type LevelStartResponse struct {
	APIStatusResponse
	Account      string `json:"account"`
	InstanceID   int    `json:"instanceId"`
	Instructions struct {
		Instructions string `json:"Instructions"`
		OrderTypes   string `json:"Order Types"`
	} `json:"instructions"`
	SecondsPerTradingDay int            `json:"secondsPerTradingDay"`
	Tickers              []string       `json:"tickers"`
	Venues               []string       `json:"venues"`
	Balances             map[string]int `json:"balances"`
}

// LevelStart starts the specified level.
func (c *Client) LevelStart(level string) (*LevelStartResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("gm/levels/%s", level)
	data, err := c.call("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &LevelStartResponse{}
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

// LevelStop stops a level.
func (c *Client) LevelStop(instance int) error {
	// Call the API.
	endpoint := fmt.Sprintf("gm/instances/%d/stop", instance)
	_, err := c.call("POST", endpoint, nil)
	if err != nil {
		return err
	}
	return nil
}
