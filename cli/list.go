package cli

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/geovanisouza92/lambdac/cli/tab"
	"github.com/geovanisouza92/lambdac/types"
)

var list = cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "List all functions",
	Action:  actionList,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "q, quiet",
			Usage: "Print only IDs",
		},
	},
}

func actionList(c *cli.Context) {
	// Get information from api
	functions, err := api.FunctionList()
	if err != nil {
		logger.Fatalf("Failed to list functions: %s", err)
	}

	if c.Bool("quiet") {
		for _, f := range functions {
			fmt.Println(f.ID)
		}
		return
	}

	// Prepare header
	t, fn := tab.New()
	defer fn()

	// Cache runtime information
	rts := map[string]types.Runtime{}

	// TODO Use channels for enhance performance

	t.Output("ID", "NAME", "RUNTIME", "MEMORY", "UPDATED")
	for _, f := range functions {
		rt, ok := rts[f.Runtime]
		if !ok {
			r, err := api.RuntimeInfo(f.Runtime)
			if err != nil {
				logger.Fatalf("Failed to get runtime information: %s", err)
			}
			rt = r
			rts[f.Runtime] = r
		}
		t.Output(f.ID[:shortIDLen], f.Name, rt.Name, f.Memory, f.Updated.Format(time.RFC822))
	}
}
