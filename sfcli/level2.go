package main

import (
	"time"

	"github.com/bearbin/go-stockfighter/sflib"
	"github.com/codegangsta/cli"
)

func level2(cli *cli.Context, api *sflib.Client) error {
	info, err := api.LevelStart("chock_a_block")
	if err != nil {
		return err
	}

	// Sleep for 5 seconds to allow the instance to warm up.
	time.Sleep(5 * time.Second)

	// Do Stuff.

	// Wait for the last order to fill.
	time.Sleep(10 * time.Second)
	// Clean everything up.
	api.LevelStop(info.InstanceID)
	return nil
}
