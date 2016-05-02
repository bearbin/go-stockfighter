package main

import (
	"fmt"
	"time"

	"github.com/bearbin/go-stockfighter/sflib"
	"github.com/codegangsta/cli"
)

func level1(cli *cli.Context, api *sflib.Client) error {
	info, err := api.LevelStart("first_steps")
	if err != nil {
		return err
	}
	// Clean everything up at the end of the function.
	defer api.LevelStop(info.InstanceID)

	// Sleep for 5 seconds to allow the instance to warm up.
	time.Sleep(5 * time.Second)

	// Send in a buy order.
	resp, err := api.StockOrder(
		info.Account,
		info.Venues[0],
		info.Tickers[0],
		0,
		100,
		"buy",
		"market",
	)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	// Wait for the final actions to complete.
	time.Sleep(10 * time.Second)
	return nil
}
