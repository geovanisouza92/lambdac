package cli

import (
	"github.com/codegangsta/cli"
	"github.com/geovanisouza92/lambdac/types"
)

var config = cli.Command{
	Name:   "config",
	Usage:  "Change function execution configuration",
	Action: actionConfig,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "f, function",
			Usage: "Function ID or name",
		},
		cli.StringFlag{
			Name:  "r, runtime",
			Usage: "Function runtime",
		},
		cli.StringFlag{
			Name:  "h, handler",
			Usage: "Function handler",
		},
		cli.StringFlag{
			Name:  "d, description",
			Usage: "Function description",
		},
		cli.StringFlag{
			Name:  "t, timeout",
			Usage: "Function execution timeout (units: ns, us (or µs), ms, s, m, h)",
			Value: "3s",
		},
		cli.IntFlag{
			Name:  "m, memory",
			Usage: "Function memory limit (in MB)",
			Value: 32,
		},
		cli.IntFlag{
			Name:  "i, instances",
			Usage: "Function max instances",
			Value: 1,
		},
	},
}

func actionConfig(c *cli.Context) {
	timeout := checkTimeoutOrFatal(c)
	function := types.Function{
		Runtime:     c.String("runtime"),
		Handler:     c.String("handler"),
		Description: c.String("description"),
		Timeout:     timeout.Nanoseconds(),
		Memory:      c.Int("memory"),
		Instances:   c.Int("instances"),
	}

	if err := api.FunctionConfig(c.String("id"), function); err != nil {
		logger.Fatalf("Error while updating function configuration", err)
	} else {
		logger.Println("Function updated.")
	}
}
