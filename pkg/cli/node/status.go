package node

import (
	api "github.com/fabiorphp/kongo"
	"github.com/fabiorphp/kongo-cli/pkg/template"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	nodeStatusTmpl = `Database:
    Reachable: {{ .Database.Reachable }}
Server:
    Total Requests: {{ .Server.TotalRequests }}
    Connections Active: {{ .Server.ConnectionsActive }}
    Connections Accepted: {{ .Server.ConnectionsAccepted }}
    Connections Handled: {{ .Server.ConnectionsHandled }}
    Connections Reading: {{ .Server.ConnectionsReading }}
    Connections Writing: {{ .Server.ConnectionsWriting }}
    Connections Waiting: {{ .Server.ConnectionsWaiting }}`

	// ErrNodeStatus retrieves an error message when client api fails.
	ErrNodeStatus = "Unable to retrieve the node status"
)

// Status retrieves usage information about a node.
func Status(c *cli.Context) error {
	client := c.App.Metadata["client"].(*api.Kongo)
	status, _, err := client.Node.Status()

	if err != nil {
		return errors.Wrap(err, ErrNodeStatus)
	}

	tmpl := template.NewPlain(nodeStatusTmpl)

	err = tmpl.Write(c.App.Writer, status)

	return err
}
