package node

import (
	"bytes"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

type (
	NodeTestSuite struct {
		suite.Suite
		assert *assert.Assertions
		buf    *bytes.Buffer
		ctx    *cli.Context
	}
)

func (s *NodeTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.buf = &bytes.Buffer{}

	app := cli.NewApp()
	app.Metadata = map[string]interface{}{}
	app.Writer = s.buf

	s.ctx = cli.NewContext(app, flag.NewFlagSet("", 0), nil)
}
