package cli

import (
	"bytes"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	buf := &bytes.Buffer{}
	set := flag.NewFlagSet("", 0)

	app := cli.NewApp()
	app.Writer = buf

	err := Version().Run(
		cli.NewContext(app, set, nil),
	)

	assert.Nil(t, err)
	assert.Contains(t, buf.String(), "Version")
	assert.Contains(t, buf.String(), "GO Version")
	assert.Contains(t, buf.String(), "OS/Arch")
}
