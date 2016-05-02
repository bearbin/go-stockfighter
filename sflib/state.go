package sflib

// State!

// Maintain a mapping of stocks to their balances, and cash balance overall.
// Maintain a list of open orders.
// Update the list of orders every few seconds. Update the stock and cash
// balances accordingly.

import (
	"sync"
	"time"
)

type stateStock struct {
	Stock
	ShareBalance int
	OutstandingBalance int
}

// State maintains and updates state.
type State struct {
	client      *Client
	account string
	Stocks      []*stateStock
	mux sync.Mutex
}

// Track adds a stock to the list of stocks to track.
func (s *State) Track(venue string, symbol string) error {
	state.Stocks = append(state.Stocks, &stateStock{Venue: venue, Symbol: string})
}

// Update updates the stocks every 5 seconds.
func (s *State) Update() {
	for {
		for i, stock := range s.Stocks {
			fmt.Println("Updating ", stock.Symbol)
			resp, err := s.client.StockOrdersStatus(stock.Venue, stock.Symbol)
			if err != nil {
				panic(err)
			}

		}
		time.Sleep(5 * time.Second)
	}
}

// NewState creates a new State object.
func NewState(c *Client, account string) *State {
	s := &State{client: c, account: account}
	go s.Update()
	return s
}
