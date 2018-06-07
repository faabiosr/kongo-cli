package services

import (
	api "github.com/fabiorphp/kongo"
	"github.com/fabiorphp/kongo-cli/pkg/template"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	tableFormat = "{{ .Id }}\t{{ .Name }}\t{{ .Protocol }}\t{{ .Host }}\t{{ .Port }}\t{{ .Path }}\n"
	tableHeader = map[string]string{
		"Id":       "ID",
		"Name":     "NAME",
		"Protocol": "PROTOCOL",
		"Host":     "HOST",
		"Port":     "PORT",
		"Path":     "PATH",
	}

	// ErrServicesList retrieves an error message when client api fails.
	ErrServicesList = "Unable to retrieve services"
)

// List retrieves a list of registred services
func List(c *cli.Context) error {
	tmpl, err := template.NewTable(tableFormat)

	if err != nil {
		return err
	}

	client := c.App.Metadata["client"].(*api.Kongo)

	services, _, err := client.Services.List(nil)

	if err != nil {
		return errors.Wrap(err, ErrServicesList)
	}

	if err := tmpl.Write(tableHeader); err != nil {
		return err
	}

	for _, svc := range services {
		if err := tmpl.Write(svc); err != nil {
			return err
		}
	}

	return tmpl.Flush(c.App.Writer)
}
