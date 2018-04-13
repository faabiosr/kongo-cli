package main

import (
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

	app.Run(os.Args)
}
