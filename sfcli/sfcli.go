package main

import (
	"os"

	"github.com/bearbin/go-stockfighter/sflib"
	"github.com/codegangsta/cli"
)

func main() {
	// Create a new API Client.
	api := sflib.NewClient("261d11b6f83a97ed6c650ace305e55a0d223645b")
	// Make sure the API is up.
	err := api.Heartbeat()
	if err != nil {
		panic(err)
	}

	// Set up the CLI information.
	app := cli.NewApp()
	app.Name = "sfcli"
	app.Usage = "complete stackfighter levels"
	app.Commands = []cli.Command{
		{
			Name:    "one",
			Aliases: []string{"1"},
			Usage:   "run the level 1 simulation",
			Action: func(c *cli.Context) {
				level1(c, api)
			},
		},
		{
			Name:    "two",
			Aliases: []string{"2"},
			Usage:   "run the level 2 simulation",
			Action: func(c *cli.Context) {
				level2(c, api)
			},
		},
	}

	// Run the application.
	app.Run(os.Args)
}

// StartAndInitialiseLevel is a derp.
func StartAndInitialiseLevel(api *sflib.Client, level string) (*sflib.LevelStartResponse, error) {
	info, err := api.LevelStart("first_steps")
	if err != nil {
		return nil, err
	}
	return info, nil
}
