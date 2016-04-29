package sflib

// State!

// Maintain a mapping of stocks to their balances, and cash balance overall.
// Maintain a list of open orders.
// Update the list of orders every few seconds. Update the stock and cash
// balances accordingly.

type stateOrder struct {
	Venue           string
	ID              int
	SharesAccounted int
}

type stateStock struct {
	Venue        string
	Symbol       string
	ShareBalance int
}

// State maintains and updates state.
type State struct {
	client      *Client
	CashBalance map[string]int
	Stocks      []stateStock
	Orders      []stateOrder
}

// BuyStock buys a stock.
func (s *State) BuyStock(
	account string,
	venue string,
	stock string,
	price int,
	quantity int,
	direction string,
	orderType string,
) (*StockOrderResponse, error) {
	// Do the request.
	resp, err := s.client.StockOrder(account, venue, stock, price, quantity, direction, orderType)
	if err != nil {
		return nil, err
	}

	// Add the returned information to the state.
	State.Orders = append(Stocks.Orders, struct{})
}

// NewState creates a new State object.
func NewState(c *Client) *State {
	return &State{client: c}
}
