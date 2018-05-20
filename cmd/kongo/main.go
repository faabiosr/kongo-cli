package main

import (
	"fmt"
	api "github.com/fabiorphp/kongo"
	"github.com/fabiorphp/kongo-cli/pkg/cli"
	ufcli "github.com/urfave/cli"
	"os"
)

var (
	appName = "kongo"
	version = "0.0.0"
)

func main() {
	app := ufcli.NewApp()
	app.Name = appName
	app.Version = version
	app.Copyright = "(c) 2018 - Fábio da Silva Ribeiro"
	app.Usage = "Manage Kong instances by CLI"
	app.Flags = []ufcli.Flag{
		ufcli.StringFlag{
			Name:   "host",
			Value:  "http://127.0.0.1:8001",
			Usage:  "Address and port of Kong instance",
			EnvVar: "KONG_HOST",
		},
	}

	app.Before = func(c *ufcli.Context) error {
		client, err := api.New(nil, c.String("host"))

		if err != nil {
			return err
		}

		c.App.Metadata["client"] = client

		return nil
	}

	app.Commands = []ufcli.Command{
		cli.Version(),
		cli.Node(),
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n\n", err)
		os.Exit(1)
	}
}
