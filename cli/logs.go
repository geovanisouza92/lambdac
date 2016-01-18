package cli

import (
	"github.com/codegangsta/cli"
)

var logs = cli.Command{
	Name:   "logs",
	Usage:  "Display logs from each function instance",
	Action: actionLogs,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, function",
			Usage:  "Function ID or Name",
			EnvVar: "LAMBDAC_FUNCTION",
		},
		cli.IntFlag{
			Name:  "t, tail",
			Usage: "Number of lines to show from the end of logs (< 1 means all lines)",
		},
	},
}

func actionLogs(c *cli.Context) {
	// TODO
}
