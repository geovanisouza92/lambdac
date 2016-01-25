package cli

import (
	"fmt"

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
	// Check required argument
	id := c.Args().First()
	if id == "" {
		id = c.String("function")
	}
	if id == "" {
		logger.Fatal("You must provide the function ID or name")
	}

	// Get runtime
	f, err := api.FunctionInfo(id)
	if err != nil {
		logger.Fatalf("Error while getting runtime information: %s", err)
	}

	r, err := api.RuntimeInfo(f.Runtime)
	if err != nil {
		logger.Fatalf("Failed to get runtime information: %s", err)
	}

	// Print information
	fmt.Printf("ID: %s\n", f.ID)
	fmt.Printf("Name: %s\n", f.Name)
	fmt.Printf("Runtime: %s\n", r.Name)
	fmt.Printf("Handler: %s\n", f.Handler)
	fmt.Printf("Description: %s\n", f.Description)
	fmt.Printf("Timeout: %v\n", f.Timeout)
	fmt.Printf("Memory: %d\n", f.Memory)
	fmt.Printf("Instances: %d\n", f.Instances)
	fmt.Println("Env:")
	for _, o := range f.Env {
		fmt.Printf("\t%s\n", o)
	}
	fmt.Printf("Created: %s\n", f.Created)
	fmt.Printf("Updated: %s\n", f.Updated)
	// TODO Should bring stats, instances, etc?
}
