package cli

import (
	"github.com/codegangsta/cli"
)

var stats = cli.Command{
	Name:   "stats",
	Usage:  "Show function statistics",
	Action: actionStats,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, function",
			Usage:  "Function ID or name",
			EnvVar: "LAMBDAC_FUNCTION",
		},
	},
}

func actionStats(c *cli.Context) {
	//
}
