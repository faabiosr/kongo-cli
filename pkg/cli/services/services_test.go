package services

import (
	"bytes"
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

type (
	ServicesTestSuite struct {
		suite.Suite
		assert *assert.Assertions
		buf    *bytes.Buffer
		ctx    *cli.Context
		flag   *flag.FlagSet
	}
)

func (s *ServicesTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.buf = &bytes.Buffer{}

	app := cli.NewApp()
	app.Metadata = map[string]interface{}{}
	app.Writer = s.buf

	s.flag = flag.NewFlagSet("test", 0)
	s.ctx = cli.NewContext(app, s.flag, nil)
}
