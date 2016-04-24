package sflib

import (
	"bytes"
	"fmt"
	"encoding/json"
)

type StockOrdersResponse struct {
	APIStatusResponse
	Venue string `json:"venue"`
	Symbol string `json:"symbol"`
	Bids []Bid `json:"bids"`
	Asks []Bid `json:"asks"`
	Timestamp string `json:"ts"`
}

func (c *Client) StockOrders(venue string, stock string) (*StockOrdersResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("venues/%s/stocks/%s", venue, stock)
	data, err := c.call("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &StockOrdersResponse{}
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

type Fill struct {
	Price int `json:"price"`
	Quantity int `json:"qty"`
	Timestamp string `json:"ts"`
}

type StockOrderResponse struct {
	APIStatusResponse
	Venue string `json:"venue"`
	Symbol string `json:"symbol"`
	Direction string `json:"direction"`
	OriginalQuantity int `json:"originalQty"`
	OutstandingQuantity int `json:"qty"`
	TotalFilled int `json:"totalFilled"`
	OrderType string `json:"orderType"`
	Price int `json:"price"`
	// An ID for this order. (Venue, ID) is guaranteed to be a unique pair.
	ID int `json:"id"`
	Account string `json:"account"`
	// The timestamp for when the order was received. (ISO-8601)
	Timestamp string `json:"ts"`
	Open bool `json:"open"`
	Fills []Fill
}

func (c *Client) StockOrder(
	account string,
	venue string,
	stock string,
	price int,
	quantity int,
	direction string,
	orderType string,
) (*StockOrderResponse, error) {
	// Encode the JSON request.
	req := &bytes.Buffer{}
	err := json.NewEncoder(req).Encode(map[string]interface{}{
		"account": account,
		"venue": venue,
		"stock": stock,
		"price": price,
		"qty": quantity,
		"direction": direction,
		"orderType": orderType,
	})

	// Call the API.
	endpoint := fmt.Sprintf("venues/%s/stocks/%s", venue, stock)
	data, err := c.call("GET", endpoint, req)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &StockOrderResponse{}
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
