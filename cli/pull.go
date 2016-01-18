package cli

import (
	"github.com/codegangsta/cli"
)

var pull = cli.Command{
	Name:      "pull",
	Usage:     "Get function code",
	ArgsUsage: "<code path (default: $CWD)>",
	Action:    actionPull,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, function",
			Usage:  "Function ID or Name",
			EnvVar: "LAMBDAC_FUNCTION",
		},
	},
}

func actionPull(c *cli.Context) {
	//
}
