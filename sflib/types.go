package sflib

// This file contains important types for decoding the JSON responses provided
// by the API.

import (
	"fmt"
)

// APIStatusResponse is a component of all API call responses, that indicates
// the status of the API or any errors that occurred during the processing of
// the API request.
type APIStatusResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

// CheckAPIStatus verifies that the API is up and the request completed successfully.
func (asr *APIStatusResponse) CheckAPIStatus() error {
	if (asr.OK != true) || (asr.Error != "") {
		return APIFailureError{asr.Error}
	}
	return nil
}

// APIFailureError is returned when the stockfighter API returns an error.
type APIFailureError struct {
	s string
}

func (afe APIFailureError) Error() string {
	return fmt.Sprintf("api returned error: %s", afe.s)
}

// Stock is the minimum representation of a stock in the stockfighter API.
type Stock struct {
	Venue  string `json:"venue"`
	Symbol string `json:"symbol"`
}

// Bid is component of the order book. It represents a bid or an ask, with a bid being "IsBuy: true"
// and an ask being "IsBuy: false".
type Bid struct {
	Price    int  `json:"price"`
	Quantity int  `json:"qty"`
	IsBuy    bool `json:"isBuy"`
}
