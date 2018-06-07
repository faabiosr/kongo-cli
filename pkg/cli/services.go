package cli

import (
	"github.com/fabiorphp/kongo-cli/pkg/cli/services"
	"github.com/urfave/cli"
)

// Services manages the Kong upstream services.
func Services() cli.Command {
	return cli.Command{
		Name:  "services",
		Usage: "Manages the Kong upstream services",
		Subcommands: []cli.Command{
			{
				Name:      "create",
				Usage:     "Creates a new service",
				ArgsUsage: "name",
				Action:    services.Create,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "protocol",
						Value: "http",
						Usage: `The protocol used to communicate with the upstream. It can be one of "http" or "https"`,
					},
					cli.StringFlag{
						Name:  "host",
						Usage: "The host of the upstream server",
					},
					cli.IntFlag{
						Name:  "port",
						Value: 80,
						Usage: "The upstream server port",
					},
					cli.StringFlag{
						Name:  "path",
						Usage: "The path to be used in requests to the upstream server",
					},
					cli.IntFlag{
						Name:  "retries",
						Value: 5,
						Usage: "The number of retries to execute upon failure the proxy",
					},
					cli.Int64Flag{
						Name:  "connect-timeout",
						Value: 60000,
						Usage: "The timeout in milliseconds for establishing a connection to the upstream server",
					},
					cli.IntFlag{
						Name:  "write-timeout",
						Value: 60000,
						Usage: "The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server",
					},
					cli.IntFlag{
						Name:  "read-timeout",
						Value: 60000,
						Usage: "The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server",
					},
					cli.StringFlag{
						Name:  "url",
						Usage: `Shorthand attribute to set "protocol", "host", "port", and "path" at once, this attribute is write-only`,
					},
				},
			},
			{
				Name:   "ls",
				Usage:  "Retrieves a list of registered services",
				Action: services.List,
			},
			{
				Name:      "rm",
				Usage:     "Delete registered services by ID or Name",
				ArgsUsage: "name [name...]",
				Action:    services.Remove,
			},
		},
	}
}
