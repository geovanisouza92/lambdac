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
		daemon,
		runtime,
	}
	app.Before = func(c *cli.Context) (err error) {
		api = client.New("http://"+c.GlobalString("host"), hc)
		cwd, err = os.Getwd()
		return
	}
	return app
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
