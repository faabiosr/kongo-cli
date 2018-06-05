package services

import (
	"fmt"
	api "github.com/fabiorphp/kongo"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	// ErrServicesCreateName retrieves an error message when name argument is not found.
	ErrServicesCreateName = errors.New(`The command requires the "name" argument`)

	// ErrServicesCreate retrieves an error message when client api fails.
	ErrServicesCreate = "Unable to create service"
)

// Create creates a new service
func Create(c *cli.Context) error {
	if c.NArg() == 0 {
		return ErrServicesCreateName
	}

	name := c.Args().Get(0)

	if name == "" {
		return ErrServicesCreateName
	}

	payload := &api.Service{
		ConnectTimeout: c.Int64("connect-timeout"),
		Host:           c.String("host"),
		Name:           name,
		Path:           c.String("path"),
		Port:           c.Int("port"),
		Protocol:       c.String("protocol"),
		ReadTimeout:    c.Int("read-timeout"),
		Retries:        c.Int("retries"),
		URL:            c.String("url"),
		WriteTimeout:   c.Int("write-timeout"),
	}

	client := c.App.Metadata["client"].(*api.Kongo)

	var svc *api.Service
	var err error

	if payload.URL != "" {
		svc, _, err = client.Services.CreateByURL(payload)
	} else {
		svc, _, err = client.Services.Create(payload)
	}

	if err != nil {
		return errors.Wrap(err, ErrServicesCreate)
	}

	fmt.Fprintln(c.App.Writer, svc.Id)

	return nil
}
