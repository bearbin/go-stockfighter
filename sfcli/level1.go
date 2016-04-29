package main

import (
	"fmt"
	"time"

	"github.com/bearbin/go-stockfighter/sflib"
	"github.com/codegangsta/cli"
)

// TODO: Work out why the first order sometimes doesn't work. Does the level have to be running frist?

func level1(cli *cli.Context, api *sflib.Client) {
	info, err := api.LevelStart("first_steps")
	if err != nil {
		panic(err)
	}

	// Sleep for 5 seconds to allow the instance to warm up.
	time.Sleep(5 * time.Second)

	// Do stuff.

	fmt.Println(response)
	// Wait for the order to fill.
	time.Sleep(10 * time.Second)
	// Clean everything up.
	api.LevelStop(info.InstanceID)
}
