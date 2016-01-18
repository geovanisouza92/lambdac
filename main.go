// LambdaC
//
// Execute code in response to events
package main // import "github.com/geovanisouza92/lambdac"

import (
	"net/http"

	"github.com/geovanisouza92/env"
	"github.com/geovanisouza92/lambdac/cli"
)

const (
	name    = "lambdac"
	usage   = "Run code in response to events"
	version = "alpha"
)

func init() {
	env.Load()
}

func main() {
	hc := &http.Client{}
	app := cli.New(name, usage, version, hc)
	app.RunAndExitOnError()
}
