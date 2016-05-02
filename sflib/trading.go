package sflib

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// HeartbeatResponse is the response to a Heartbeat API call.
// It simply informs the user if the API is up and functioning correctly.
type HeartbeatResponse struct {
	APIStatusResponse
}

// Heartbeat checks to make sure that the API is up and operational.
// If the API is not functioning correctly an informative APIFailureError will
// be returned.
func (c *Client) Heartbeat() error {
	// Call the API.
	data, err := c.call("GET", "ob/api/heartbeat", nil)
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

// VenueHeartbeatResponse is the responsefor the VenueHeartbeat method.
// In addition to the APIStatusResponse it only contains the venue which the
// check was performed on.
type VenueHeartbeatResponse struct {
	APIStatusResponse
	Venue string `json:"venue"`
}

// VenueHeartbeat checks if a venue is up and running.
func (c *Client) VenueHeartbeat(venue string) (*VenueHeartbeatResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("ob/api/venues/%s/heartbeat", venue)
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

// VenueSymbol represents a stock symbol and the human-readable name.
type VenueSymbol struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// VenueStocksResponse is the response from the SF API for the VenueStocks method.
type VenueStocksResponse struct {
	APIStatusResponse
	Symbols []VenueSymbol `json:"symbols"`
}

// VenueStocks returns a list of stocks on the specified exchange.
func (c *Client) VenueStocks(venue string) (*VenueStocksResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("ob/api/venues/%s/stocks", venue)
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

// StockOrdersResponse is the response from the StockOrders method of the SF API.
type StockOrdersResponse struct {
	APIStatusResponse
	Stock
	Bids      []Bid  `json:"bids"`
	Asks      []Bid  `json:"asks"`
	Timestamp string `json:"ts"`
}

// StockOrders gets the orders for the specified stock on the specified exchange (venue).
func (c *Client) StockOrders(venue string, symbol string) (*StockOrdersResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("/ob/api/venues/%s/stocks/%s", venue, symbol)
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

// Fill represents a transaction that has taken place.
type Fill struct {
	Price     int    `json:"price"`
	Quantity  int    `json:"qty"`
	Timestamp string `json:"ts"`
}

// StockOrderResponse is the response from the StockOrder method of the SF API.
type StockOrderResponse struct {
	APIStatusResponse
	Stock
	Direction           string `json:"direction"`
	OriginalQuantity    int    `json:"originalQty"`
	OutstandingQuantity int    `json:"qty"`
	TotalFilled         int    `json:"totalFilled"`
	OrderType           string `json:"orderType"`
	Price               int    `json:"price"`
	// An ID for this order. (Venue, ID) is guaranteed to be a unique pair.
	ID      int    `json:"id"`
	Account string `json:"account"`
	// The timestamp for when the order was received. (ISO-8601)
	Timestamp string `json:"ts"`
	Open      bool   `json:"open"`
	Fills     []Fill
}

// StockOrder calls the StockOrder method of the SF API and returns the decoded result.
func (c *Client) StockOrder(
	account string,
	venue string,
	symbol string,
	price int,
	quantity int,
	direction string,
	orderType string,
) (*StockOrderResponse, error) {
	// Encode the JSON request.
	req := &bytes.Buffer{}
	err := json.NewEncoder(req).Encode(map[string]interface{}{
		"account":   account,
		"venue":     venue,
		"stock":     symbol,
		"price":     price,
		"qty":       quantity,
		"direction": direction,
		"orderType": orderType,
	})

	// Call the API.
	endpoint := fmt.Sprintf("ob/api/venues/%s/stocks/%s/orders", venue, symbol)
	data, err := c.call("POST", endpoint, req)
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

// StockQuoteResponse is the response returned by the stockfighter API after a
// call to the StockQuote method.
type StockQuoteResponse struct {
	APIStatusResponse
	Stock
	// Best current bid for the stock.
	Bid int `json:"bid"`
	// Best current ask for the stock.
	Ask int `json:"ask"`
	// Total size of orders at the best bid.
	BidSize int `json:"bidSize"`
	// Total size of orders at the best ask.
	AskSize int `json:"askSize"`
	// Total size of bids at any price.
	BidDepth int `json:"bidDepth"`
	// Total size of asks at any price.
	AskDepth int `json:"askDepth"`
	// Price of the last trade.
	LastTradePrice int `json:"last"`
	// Quantity of the last trade.
	LastTradeQuantity int `json:"lastSize"`
	// ISO-8601 timestamp of the last trade.
	LastTradeTime string `json:"lastTrade"`
	// ISO-8601 timestamp of the last quote.
	QuoteTimestamp string `json:"quoteTime"`
}

// StockQuote gets a quote for the specified stock on the specified excnange.
func (c *Client) StockQuote(venue string, symbol string) (*StockQuoteResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("/ob/api/venues/%s/stocks/%s/quote", venue, symbol)
	data, err := c.call("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &StockQuoteResponse{}
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

// Type OrderStatus is the status of an order.
type OrderStatus struct {
	Stock
	// Buy or Sell
	Direction           string `json:"direction"`
	OriginalQuantity    int    `json:"originalQty"`
	OutstandingQuantity int    `json:"qty"`
	Price               int    `json:"price"`
	OrderType           string `json:"orderType"`
	ID                  int    `json:"id"`
	Account             string `json:"account"`
	// ISO-8601 timestamp of the status response.
	Timestamp   string `json:"ts"`
	TotalFilled int    `json:"totalFilled"`
	Open        bool   `json:"open"`
	Fills       []Fill
}

// OrderStatusResponse is the response returned by the StockFighter API giving
// the status of an order.
type OrderStatusResponse struct {
	APIStatusResponse
	OrderStatus
}

// OrderStatus gets the status for a given order in a given stock on the
// specified exchange.
func (c *Client) OrderStatus(venue string, symbol string, id int) (*OrderStatusResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("/ob/api/venues/%s/stocks/%s/orders/%d", venue, symbol, id)
	data, err := c.call("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &OrderStatusResponse{}
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

// OrderCancel cancels the given order, then returns the status of the order.
func (c *Client) OrderCancel(venue string, symbol string, id int) (*OrderStatusResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("/ob/api/venues/%s/stocks/%s/orders/%d", venue, symbol, id)
	data, err := c.call("DELETE", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &OrderStatusResponse{}
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

// VenueOrdersResponse is the response to a VenueOrders call.
type VenueOrdersResponse struct {
	APIStatusResponse
	Venue string `json:"venue"`
	Orders []OrderStatus `json:"orders"`
}

// VenueOrdersStatus gets the status for all orders on the specified exchange.
func (c *Client) VenueOrdersStatus(venue string, account string) (*VenueOrdersResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("/ob/api/venues/%s/accounts/%s/orders", venue, account)
	data, err := c.call("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &VenueOrdersResponse{}
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

// StockOrdersStatus gets the status for all orders on the specified exchange/stock.
func (c *Client) StockOrdersStatus(venue string, account string, symbol string) (*VenueOrdersResponse, error) {
	// Call the API.
	endpoint := fmt.Sprintf("/ob/api/venues/%s/accounts/%s/stocks/%s/orders", venue, account, symbol)
	data, err := c.call("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response.
	response := &VenueOrdersResponse{}
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
