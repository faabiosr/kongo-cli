package services

import (
	"fmt"
	api "github.com/fabiorphp/kongo"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	// ErrServicesRemoveName retrieves an error message when name argument is not found.
	ErrServicesRemoveName = errors.New(`The command requires at least one argument`)

	// ErrServicesRemoveStatus retrieves an error message when one of deletions failed.
	ErrServicesRemoveStatus = errors.New(`One of deletions failed`)

	// ErrServicesRemove retrieves an error message when client api fails.
	ErrServicesRemove = "Unable to remove service"
)

// Remove deletes registred service by ID or Name
func Remove(c *cli.Context) error {
	if c.NArg() == 0 {
		return ErrServicesRemoveName
	}

	client := c.App.Metadata["client"].(*api.Kongo)
	failed := false

	for _, name := range c.Args() {
		if _, err := client.Services.Delete(name); err != nil {
			fmt.Fprintln(c.App.Writer, errors.Wrap(err, ErrServicesRemove))

			failed = true
			continue
		}

		fmt.Fprintln(c.App.Writer, name)
	}

	if failed {
		return ErrServicesRemoveStatus
	}

	return nil
}
