package cli

import (
	"github.com/codegangsta/cli"
)

var destroy = cli.Command{
	Name:    "destroy",
	Aliases: []string{"rm"},
	Usage:   "Destroy a function, its code and logs",
	Action:  actionDestroy,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, function",
			Usage:  "Function ID or name",
			EnvVar: "LAMBDAC_FUNCTION",
		},
	},
}

func actionDestroy(c *cli.Context) {
	// TODO
}
