package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/codegangsta/cli"
	"github.com/geovanisouza92/lambdac/cli/tab"
	"github.com/geovanisouza92/lambdac/types"
)

var runtime = cli.Command{
	Name:    "runtime",
	Aliases: []string{"rt"},
	Usage:   "Function runtimes information",
	Action:  actionRuntime,
	Subcommands: []cli.Command{
		cli.Command{
			Name:   "create",
			Usage:  "Create a new runtime",
			Action: actionRuntimeCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t, template",
					Usage: "Template file to create multiple runtimes at once (override all other flags)",
				},
				cli.StringFlag{
					Name:  "n, name",
					Usage: "Runtime name",
				},
				cli.StringFlag{
					Name:  "l, label",
					Usage: "Runtime label",
				},
				cli.StringFlag{
					Name:  "i, image",
					Usage: "Runtime image",
				},
				cli.StringFlag{
					Name:  "c, command",
					Usage: "Runtime command",
				},
				cli.BoolFlag{
					Name:  "a, agent",
					Usage: "Run functions in this runtime as daemon, processing multiple events with lower latency",
				},
				cli.StringFlag{
					Name:  "d, driver",
					Usage: "Runtime backend driver to create instances",
					Value: "docker",
				},
				cli.StringSliceFlag{
					Name:  "o, driver-opt",
					Usage: "Runtime driver options used to create instances",
				},
			},
		},
		cli.Command{
			Name:      "info",
			Usage:     "Display runtime information",
			ArgsUsage: "Runtime ID or name",
			Action:    actionRuntimeInfo,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "r, runtime",
					Usage: "Runtime ID or name",
				},
			},
		},
		cli.Command{
			Name:      "destroy",
			Aliases:   []string{"rm"},
			Usage:     "Destroy a runtime",
			ArgsUsage: "Runtime ID or name",
			Action:    actionRuntimeDestroy,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "r, runtime",
					Usage: "Runtime ID or name",
				},
				cli.BoolFlag{
					Name:  "f, force",
					Usage: "Destroy runtime and functions at once",
				},
			},
		},
	},
}

type templateConfig struct {
	Runtimes types.Runtimes `json:"runtimes"`
}

func actionRuntime(c *cli.Context) {
	// Call API
	runtimes, err := api.RuntimeList()
	if err != nil {
		logger.Fatalf("Failed to list runtimes: %s", err)
	}

	// Print information
	t, fn := tab.New()
	defer fn()

	t.Output("ID", "NAME", "LABEL", "IMAGE", "DRIVER", "UPDATED")
	for _, r := range runtimes {
		t.Output(r.ID[:shortIDLen], r.Name, r.Label, r.Image, r.Driver, r.Updated.Format(time.RFC822))
	}
}

func actionRuntimeCreate(c *cli.Context) {
	// Create from template
	tplName := c.String("template")
	if tplName != "" {
		runtimeCreateFromTemplate(c, tplName)
		return
	}

	// Create from flags
	r := types.Runtime{
		Name:    c.String("name"),
		Label:   c.String("label"),
		Image:   c.String("image"),
		Command: c.String("command"),
		Agent:   c.Bool("agent"),
		Driver:  c.String("driver"),
		Options: c.StringSlice("driver-opt"),
	}

	runtimeCreate(r)
}

func runtimeCreateFromTemplate(c *cli.Context, tplName string) {
	// Read template
	data, err := ioutil.ReadFile(tplName)
	if err != nil {
		logger.Fatalf("Error while reading template file: %s", err)
	}

	// Parse template
	var tpl templateConfig
	if err = json.Unmarshal(data, &tpl); err != nil {
		logger.Fatalf("Error while parsing template file: %s", err)
	}

	// Create each runtime
	for _, r := range tpl.Runtimes {
		if r.Driver == "" && c.String("driver") != "" {
			r.Driver = c.String("driver")
		}
		if len(r.Options) == 0 && len(c.StringSlice("driver-opt")) > 0 {
			r.Options = c.StringSlice("driver-opt")
		}
		runtimeCreate(r)
	}
	logger.Println("All runtimes was created.")
}

func runtimeCreate(r types.Runtime) {
	// Check required attributes
	requiredFlags := map[string]string{
		"name":    r.Name,
		"image":   r.Image,
		"command": r.Command,
		"driver":  r.Driver,
	}
	for f, v := range requiredFlags {
		if v == "" {
			logger.Fatalf("%q is required", f)
		}
	}

	// Call API
	created, err := api.RuntimeCreate(r)
	if err != nil {
		logger.Fatalf("Error while creating runtime: %s", err)
	} else {
		logger.Printf("Runtime %q (%s) created.\n", created.Name, created.ID[:shortIDLen])
	}
}

func actionRuntimeInfo(c *cli.Context) {
	// Check required argument
	id := c.Args().First()
	if id == "" {
		id = c.String("runtime")
	}
	if id == "" {
		logger.Fatal("You must provide the runtime ID or name")
	}

	// Get runtime
	r, err := api.RuntimeInfo(id)
	if err != nil {
		logger.Fatalf("Error while getting runtime information: %s", err)
	}

	// Print information
	fmt.Printf("ID: %s\n", r.ID)
	fmt.Printf("Name: %s\n", r.Name)
	fmt.Printf("Label: %s\n", r.Label)
	fmt.Printf("Image: %s\n", r.Image)
	fmt.Printf("Command: %s\n", r.Command)
	fmt.Printf("Agent: %v\n", r.Agent)
	fmt.Printf("Driver: %s\n", r.Driver)
	fmt.Println("Options:")
	for _, o := range r.Options {
		fmt.Printf("\t%s\n", o)
	}
	fmt.Printf("Created: %s\n", r.Created)
	fmt.Printf("Updated: %s\n", r.Updated)
}

func actionRuntimeDestroy(c *cli.Context) {
	// Check required argument
	id := c.Args().First()
	if id == "" {
		id = c.String("runtime")
	}
	if id == "" {
		logger.Fatal("You must provide the runtime ID or name")
	}

	// Get runtime
	r, err := api.RuntimeInfo(id)
	if err != nil {
		logger.Fatalf("Error while getting runtime information: %s", err)
	}

	// Prompt user if --force is not set
	if !c.Bool("force") && !promptYesNo("Are you sure you want to delete the runtime %q (ID: %s)?", r.Label, r.ID) {
		return
	}

	// Delete runtime
}
