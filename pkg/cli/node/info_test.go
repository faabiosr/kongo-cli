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
	InfoTestSuite struct {
		suite.Suite
		assert *assert.Assertions
		buf    *bytes.Buffer
		ctx    *cli.Context
	}
)

func (s *InfoTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.buf = &bytes.Buffer{}

	app := cli.NewApp()
	app.Metadata = map[string]interface{}{}
	app.Writer = s.buf

	s.ctx = cli.NewContext(app, flag.NewFlagSet("", 0), nil)
}

func (s *InfoTestSuite) TestInfoWhenApiReturnsError() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Node: &MockNode{Error: true}}

	err := Info(s.ctx)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrNodeInfo)
}

func (s *InfoTestSuite) TestInfo() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Node: &MockNode{}}

	err := Info(s.ctx)
	res := s.buf.String()

	s.assert.Nil(err)
	s.assert.Contains(res, "Configuration")
	s.assert.Contains(res, "Hostname")
	s.assert.Contains(res, "Lua version")
	s.assert.Contains(res, "Plugins")
	s.assert.Contains(res, "Prng seeds")
	s.assert.Contains(res, "Tagline")
	s.assert.Contains(res, "Timers")
	s.assert.Contains(res, "Version")
}

func (s *InfoTestSuite) TestInfoTemplateError() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Node: &MockNode{EmptyStatus: true}}

	err := Info(s.ctx)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), template.ErrTemplateParsing)
}

func TestInfoTestSuite(t *testing.T) {
	suite.Run(t, new(InfoTestSuite))
}
