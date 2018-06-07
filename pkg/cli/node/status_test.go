package node

import (
	api "github.com/fabiorphp/kongo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	StatusTestSuite struct {
		NodeTestSuite
	}
)

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

func TestStatusTestSuite(t *testing.T) {
	suite.Run(t, new(StatusTestSuite))
}
