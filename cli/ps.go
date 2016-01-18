package cli

import (
	"github.com/codegangsta/cli"
)

var ps = cli.Command{
	Name:   "ps",
	Usage:  "List function instances",
	Action: actionPs,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, function",
			Usage:  "Function ID or Name",
			EnvVar: "LAMBDAC_FUNCTION",
		},
	},
}

func actionPs(c *cli.Context) {
	// TODO
}
