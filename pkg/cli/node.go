package cli

import (
	"github.com/fabiorphp/kongo-cli/pkg/cli/node"
	"github.com/urfave/cli"
)

// Node retrives details about a node
func Node() cli.Command {
	return cli.Command{
		Name:  "node",
		Usage: "Retrieves details about a node",
		Subcommands: []cli.Command{
			cli.Command{
				Name:   "status",
				Usage:  "Retrieves usage information about a node",
				Action: node.Status,
			},
		},
	}
}
