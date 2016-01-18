package cli

import (
	"github.com/codegangsta/cli"
)

var push = cli.Command{
	Name:      "push",
	Usage:     "Update function code",
	ArgsUsage: "<code path (default: $CWD)>",
	Action:    actionPush,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, function",
			Usage:  "Function ID or Name",
			EnvVar: "LAMBDAC_FUNCTION",
		},
	},
}

func actionPush(c *cli.Context) {
	// TODO
}
