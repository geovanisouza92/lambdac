package cli

import (
	"github.com/codegangsta/cli"
	"github.com/geovanisouza92/lambdac/types"
)

var create = cli.Command{
	Name:      "create",
	Usage:     "Create a new function",
	ArgsUsage: "<code path>",
	Action:    actionCreate,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "n, name",
			Usage: "Function name",
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
		cli.StringSliceFlag{
			Name:  "e, env",
			Usage: "Function environment variables (format: VARIABLE=VALUE ; can be informed multiple times)",
		},
	},
}

func actionCreate(c *cli.Context) {
	codePath := c.Args().First()
	if codePath == "" {
		codePath = cwd
	}
	if codePath == "" {
		// TODO Warning?
	}
	codePath = codePath

	// Check required params
	requiredAttrs := map[string]string{
		"name":    c.String("name"),
		"runtime": c.String("runtime"),
	}
	for a, v := range requiredAttrs {
		if v == "" {
			logger.Fatalf("%q is required", a)
		}
	}

	timeout := checkTimeoutOrFatal(c)
	function := types.Function{
		Name:        c.String("name"),
		Runtime:     c.String("runtime"),
		Handler:     c.String("handler"),
		Description: c.String("description"),
		Timeout:     timeout.Nanoseconds(),
		Memory:      c.Int("memory"),
		Instances:   c.Int("instances"),
		Env:         c.StringSlice("env"),
	}

	created, err := api.FunctionCreate(function)
	if err != nil {
		logger.Fatalf("Error while creating function: %s", err)
	} else {
		logger.Printf("Function %q (%q) created.\n", created.Name, created.ID[:shortIDLen])
	}
}
