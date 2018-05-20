package node

import (
	"bytes"
	"flag"
	api "github.com/fabiorphp/kongo"
	"github.com/fabiorphp/kongo-cli/pkg/template"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
	"testing"
)

type (
	StatusTestSuite struct {
		suite.Suite
		assert *assert.Assertions
		buf    *bytes.Buffer
		ctx    *cli.Context
	}
)

func (s *StatusTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.buf = &bytes.Buffer{}

	app := cli.NewApp()
	app.Metadata = map[string]interface{}{}
	app.Writer = s.buf

	s.ctx = cli.NewContext(app, flag.NewFlagSet("", 0), nil)
}

func (s *StatusTestSuite) TestStatusWhenApiReturnsError() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Node: &MockNode{Error: true}}

	err := Status(s.ctx)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrNodeStatus)
}

func (s *StatusTestSuite) TestStatus() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Node: &MockNode{}}

	err := Status(s.ctx)
	res := s.buf.String()

	s.assert.Nil(err)
	s.assert.Contains(res, "Database")
	s.assert.Contains(res, "Reachable")
	s.assert.Contains(res, "Server")
	s.assert.Contains(res, "Total Requests")
	s.assert.Contains(res, "Connections")
	s.assert.Contains(res, "Active")
	s.assert.Contains(res, "Accepted")
	s.assert.Contains(res, "Handled")
	s.assert.Contains(res, "Reading")
	s.assert.Contains(res, "Writing")
	s.assert.Contains(res, "Waiting")
}

func (s *StatusTestSuite) TestNodeStatusCommandTemplateError() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Node: &MockNode{EmptyStatus: true}}

	err := Status(s.ctx)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), template.ErrTemplateParsing)
}

func TestStatusTestSuite(t *testing.T) {
	suite.Run(t, new(StatusTestSuite))
}
