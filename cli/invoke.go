package cli

import (
	"github.com/codegangsta/cli"
)

var invoke = cli.Command{
	Name:      "invoke",
	Usage:     "Invoke function",
	ArgsUsage: "<event data>",
	Action:    actionInvoke,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, function",
			Usage:  "Function ID or name",
			EnvVar: "LAMBDAC_FUNCTION",
		},
	},
}

func actionInvoke(c *cli.Context) {
	// TODO Use arguments
}
