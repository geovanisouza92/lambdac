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
		cli.BoolFlag{
			Name:  "F, force",
			Usage: "Destroy function without asking",
		},
	},
}

func actionDestroy(c *cli.Context) {
	// Check required argument
	id := c.Args().First()
	if id == "" {
		id = c.String("function")
	}
	if id == "" {
		logger.Fatal("You must provide the function ID or name")
	}

	// Get function
	f, err := api.FunctionInfo(id)
	if err != nil {
		logger.Fatalf("Error while getting function information: %s", err)
	}

	// Prompt user if --force is not set
	if !c.Bool("force") {
		if !promptYesNo("Are you sure you want to delete the function %q (ID: %s)?", f.Name, f.ID) {
			return
		}
	}

	// Delete function
	if err = api.FunctionDestroy(f.ID, c.Bool("force")); err != nil {
		logger.Fatalf("Error while destroying function: %s", err)
	}
}
