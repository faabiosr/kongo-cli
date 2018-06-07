package services

import (
	api "github.com/fabiorphp/kongo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	ListTestSuite struct {
		ServicesTestSuite
	}
)

func (s *ListTestSuite) TestRetrievesErrorWhenApiFails() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{Error: true}}

	err := List(s.ctx)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrServicesList)
}

func (s *ListTestSuite) TestCreate() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{}}

	err := List(s.ctx)
	res := s.buf.String()

	s.assert.Nil(err)
	s.assert.Contains(res, "a1")
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}
