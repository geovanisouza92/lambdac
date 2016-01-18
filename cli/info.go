package cli

import (
	"github.com/codegangsta/cli"
)

var info = cli.Command{
	Name:   "info",
	Usage:  "Get detailed information about a function",
	Action: actionInfo,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, function",
			Usage:  "Function ID or name",
			EnvVar: "LAMBDAC_FUNCTION",
		},
	},
}

func actionInfo(c *cli.Context) {
	// TODO
}
