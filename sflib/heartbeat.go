package sflib

import (
	"encoding/json"
)

// HeartbeatResponse is the response to a heartbeat API call.
type HeartbeatResponse struct {
	APIStatusResponse
}

// Heartbeat checks to make sure that the API is up and operational.
// If the API is up, true will be returned.
func (c *Client) Heartbeat() error {
	// Call the API.
	data, err := c.call("GET", "heartbeat", nil)
	if err != nil {
		return err
	}

	// Unmarshal the JSON response.
	response := &HeartbeatResponse{}
	err = json.NewDecoder(*data).Decode(response)
	if err != nil {
		return err
	}

	// Check if the request completed successfully.
	err = response.CheckAPIStatus()
	if err != nil {
		return err
	}

	// If nothing failed, then the API must be working correctly.
	return nil
}
