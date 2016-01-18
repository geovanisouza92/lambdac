package cli

import (
	"time"

	"github.com/codegangsta/cli"
	"github.com/geovanisouza92/lambdac/cli/tab"
)

var list = cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "List all functions",
	Action:  actionList,
}

func actionList(c *cli.Context) {
	// Get information from api
	functions, err := api.FunctionList()
	if err != nil {
		logger.Fatalf("Failed to list functions: %s", err)
	}

	// Prepare header
	t, fn := tab.New()
	defer fn()

	t.Output("ID", "NAME", "RUNTIME", "MEMORY", "UPDATED")
	for _, f := range functions {
		t.Output(f.ID[:shortIDLen], f.Name, f.Runtime, f.Memory, f.Updated.Format(time.RFC822))
	}
}
