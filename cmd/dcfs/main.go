package main

import (
	"log"
	"os"

	"github.com/c0rby/dcfs/internal/actions"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "dcfs",
		Usage: "A tool for working with the decomposed fs",
		Commands: []*cli.Command{
			{
				Name:      "tree",
				Aliases:   []string{"tr"},
				ArgsUsage: "[path]",
				Usage:     "List the contents of the decomposed fs in a tree-like format",
				Action: func(c *cli.Context) error {
					path := c.Args().First()
					return actions.Tree(path)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
