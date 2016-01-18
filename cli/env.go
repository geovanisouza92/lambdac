package cli

import (
	"github.com/codegangsta/cli"
	envLib "github.com/geovanisouza92/env"
	"github.com/geovanisouza92/lambdac/cli/tab"
)

var env = cli.Command{
	Name:   "env",
	Usage:  "Function environment configuration",
	Action: actionEnv,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, function",
			Usage:  "Function ID or name",
			EnvVar: "LAMBDAC_FUNCTION",
		},
	},
	Subcommands: []cli.Command{
		cli.Command{
			Name:      "set",
			Usage:     "Set variable value",
			ArgsUsage: "VARIABLE=VALUE VARIABLE=VALUE ...",
			Action:    actionEnvSet,
		},
		cli.Command{
			Name:      "unset",
			Usage:     "Unset variable value",
			ArgsUsage: "VARIABLE VARIABLE ...",
			Action:    actionEnvUnset,
		},
	},
}

func actionEnv(c *cli.Context) {
	fid := checkFidOrFatal(c)

	envVars, err := api.FunctionEnv(fid)
	if err != nil {
		logger.Fatalf("Failed to list runtimes: %s", err)
	}

	t, fn := tab.New()
	defer fn()

	for _, e := range envVars {
		t.Output(e)
	}
}

func actionEnvSet(c *cli.Context) {
	fid := checkFidOrFatal(c)

	if !c.Args().Present() {
		// Inform user to type vars
		logger.Println("You must inform environment variables in format VAR=VALUE")
		return
	}

	vars := make([]string, len(c.Args()))
	for _, a := range c.Args() {
		if _, _, err := envLib.ParseLine(a); err != nil {
			logger.Fatalf("Invalid environment variable %q. Pattern must be VARIABLE=VALUE", err)
		}
		vars = append(vars, a)
	}

	if err := api.FunctionEnvSet(fid, vars); err != nil {
		logger.Fatalf("Error while settings environment variables: %s", err)
	} else {
		logger.Println("Environment variables was set.")
	}
}

func actionEnvUnset(c *cli.Context) {
	// fid := checkFidOrFatal(c)
	// TODO
}
