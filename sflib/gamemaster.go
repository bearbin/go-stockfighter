package sflib

import (
	"encoding/json"
	"fmt"
)

// LevelStartResponse is a response from the gamemaster API.
type LevelStartResponse struct {
	APIStatusResponse
	Account              string            `json:"account"`
	InstanceID           int               `json:"instanceId"`
	Instructions         map[string]string `json:"instructions"`
	SecondsPerTradingDay int               `json:"secondsPerTradingDay"`
	Tickers              []string          `json:"tickers"`
	Venues               []string          `json:"venues"`
	Balances             map[string]int    `json:"balances"`
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

// LevelRestart restarts a level, keeping everything the same.
func (c *Client) LevelRestart(instance int) error {
	// Call the API.
	endpoint := fmt.Sprintf("gm/instances/%d/restart", instance)
	_, err := c.call("POST", endpoint, nil)
	if err != nil {
		return err
	}
	return nil
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

// LevelResume resumes a level. Given an instanceID, it returns a LevelStartResponse.
func (c *Client) LevelResume(instance int) (*LevelStartResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("gm/instances/%d/resume", instance)
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

// LevelGetResponse is the response to a level get command.
type LevelGetResponse struct {
	APIStatusResponse
	Details struct {
		EndOfTheWorldDay int `json:"endOfTheWorldDay"`
		TradingDay       int `json:"tradingDay"`
	} `json:"details"`
	Done       bool   `json:"done"`
	InstanceID int    `json:"id"`
	State      string `json:"state"`
}

// LevelGet gets level information given the instance ID.
func (c *Client) LevelGet(instance int) (*LevelGetResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("gm/instances/%d", instance)
	data, err := c.call("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &LevelGetResponse{}
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
