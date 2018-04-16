package cli

import (
	"flag"
	"github.com/stretchr/testify/suite"
	"github.com/urfave/cli"
	"testing"
)

type (
	VersionTestSuite struct {
		CommandTestSuite
	}
)

func (s *VersionTestSuite) TestVersionCommand() {
	set := flag.NewFlagSet("", 0)
	ctx := cli.NewContext(s.app, set, nil)

	err := Version().Run(ctx)

	s.assert.Nil(err)
	s.assert.Contains(s.buf.String(), "Version")
	s.assert.Contains(s.buf.String(), "GO Version")
	s.assert.Contains(s.buf.String(), "OS/Arch")
}

func TestVersionTestSuite(t *testing.T) {
	suite.Run(t, new(VersionTestSuite))
}
