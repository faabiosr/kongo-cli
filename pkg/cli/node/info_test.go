package node

import (
	api "github.com/fabiorphp/kongo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	InfoTestSuite struct {
		NodeTestSuite
	}
)

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

func TestInfoTestSuite(t *testing.T) {
	suite.Run(t, new(InfoTestSuite))
}
