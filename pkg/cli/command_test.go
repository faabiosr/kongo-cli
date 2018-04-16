package cli

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
)

type (
	CommandTestSuite struct {
		suite.Suite

		assert *assert.Assertions
		app    *cli.App
		buf    *bytes.Buffer
	}
)

func (s *CommandTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
	s.app = cli.NewApp()
	s.buf = &bytes.Buffer{}

	s.app.Writer = s.buf
}
