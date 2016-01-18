// Command-line interface
package cli

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/geovanisouza92/lambdac/client"
)

const shortIDLen = 8

var (
	logger *log.Logger
	api    client.API
	cwd    string
)

func init() {
	logger = log.New(os.Stdout, "[cli] ", 0)
}

func New(name, usage, version string, hc *http.Client) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "H, host",
			Usage:  "LambdaC host",
			EnvVar: "LAMBDAC_HOST",
			Value:  "localhost",
		},
	}
	app.Commands = []cli.Command{
		config,
		create,
		daemon,
		destroy,
		env,
		info,
		invoke,
		list,
		logs,
		ps,
		pull,
		push,
		runtime,
		stats,
	}
	app.Before = func(c *cli.Context) (err error) {
		api = client.New("http://"+c.GlobalString("host"), hc)
		cwd, err = os.Getwd()
		return
	}
	return app
}

func checkFidOrFatal(c *cli.Context) string {
	fid := c.String("function")
	if fid == "" {
		logger.Fatal("You must inform function ID")
	}
	return fid
}

func checkTimeoutOrFatal(c *cli.Context) time.Duration {
	timeout, err := time.ParseDuration(c.String("timeout"))
	if err != nil {
		logger.Fatalf("invalid timeout: %s", err)
	}
	return timeout
}

func promptYesNo(msg string, v ...interface{}) (result bool) {
	fmt.Printf(msg, v...)
	fmt.Print(" (yes/no): ")
	for {
		var answer string
		fmt.Scanln(&answer)
		switch answer {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Print("Please type 'yes' or 'no': ")
		}
	}
}
