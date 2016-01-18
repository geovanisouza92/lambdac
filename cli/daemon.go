package cli

import (
	"time"

	"github.com/codegangsta/cli"
	"github.com/geovanisouza92/lambdac/server"
	"gopkg.in/tylerb/graceful.v1"
)

var daemon = cli.Command{
	Name:      "daemon",
	Usage:     "Start daemon service",
	Action:    actionDaemon,
	ArgsUsage: "<IP:HOST>",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "s, store",
			Usage:  "Store driver",
			EnvVar: "LAMBDAC_STORE",
			Value:  "memory",
		},
		cli.StringFlag{
			Name:   "c, conn",
			Usage:  "Connection string to store",
			EnvVar: "LAMBDAC_CONN",
		},
	},
}

func actionDaemon(c *cli.Context) {
	// Get params
	store := c.String("store")
	conn := c.String("conn")
	host := c.Args().First()
	if host == "" {
		host = c.GlobalString("host")
	}

	logger.Printf("Using %q store driver.", store)

	// Create server
	s, err := server.New(store, conn)
	if err != nil {
		logger.Fatalf("Failed to start daemon: %s", err)
	}
	defer s.Close()

	// Start server
	logger.Println("Ready.")
	graceful.Run(host, 5*time.Second, s)
}
