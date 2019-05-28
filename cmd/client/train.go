package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var trainCommand = cli.Command{
	Name:  "train",
	Usage: "query/add/remove an train",
	Subcommands: []cli.Command{
		{
			Name:  "list",
			Usage: "list the trains",
			Action: func(c *cli.Context) error {
				fmt.Println("list train")
				return nil
			},
		},
	},
}
