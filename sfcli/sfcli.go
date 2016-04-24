package main

import (
	"github.com/bearbin/go-stockfighter/sflib"
	"fmt"
)

func main() {
	cli := sflib.NewClient("")
	err := cli.Heartbeat()
	if err != nil {
		panic(err)
	}
	fmt.Println("API UP")
	_, err = cli.VenueHeartbeat("TESTEX")
	if err != nil {
		panic(err)
	}
	fmt.Println("TESTEX UP")
	stocks, err := cli.VenueStocks("TESTEX")
	if err != nil {
		panic(err)
	}
	fmt.Println(stocks.Symbols)
	stock, err := cli.StockOrders("TESTEX", stocks.Symbols[0].Symbol)
	if err != nil {
		panic(err)
	}
	fmt.Println(stock.Timestamp)
}
