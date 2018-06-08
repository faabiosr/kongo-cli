package services

import (
	"fmt"
	api "github.com/fabiorphp/kongo"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	// ErrServicesUpdateID retrieves an error message when id argument is not found.
	ErrServicesUpdateID = errors.New(`The command requires the "id" or "name" argument`)

	// ErrServicesUpdate retrieves an error message when client api fails.
	ErrServicesUpdate = "Unable to update service"
)

// Update creates a new service
func Update(c *cli.Context) error {
	if c.NArg() == 0 {
		return ErrServicesUpdateID
	}

	id := c.Args().Get(0)

	if id == "" {
		return ErrServicesUpdateID
	}

	client := c.App.Metadata["client"].(*api.Kongo)

	svc, _, err := client.Services.Get(id)

	if err != nil {
		return errors.Wrap(err, ErrServicesUpdate)
	}

	svc.ConnectTimeout = c.Int64("connect-timeout")
	svc.Path = c.String("path")
	svc.Port = c.Int("port")
	svc.ReadTimeout = c.Int("read-timeout")
	svc.Retries = c.Int("retries")
	svc.WriteTimeout = c.Int("write-timeout")
	svc.URL = c.String("url")

	if v := c.String("host"); v != "" {
		svc.Host = v
	}

	if v := c.String("name"); v != "" {
		svc.Name = v
	}

	if v := c.String("protocol"); v != "" {
		svc.Protocol = v
	}

	if svc.URL != "" {
		svc, _, err = client.Services.UpdateByURL(id, svc)
	} else {
		svc, _, err = client.Services.Update(id, svc)
	}

	if err != nil {
		return errors.Wrap(err, ErrServicesUpdate)
	}

	fmt.Fprintln(c.App.Writer, svc.Id)

	return nil
}
