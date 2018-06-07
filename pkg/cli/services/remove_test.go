package services

import (
	api "github.com/fabiorphp/kongo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	RemoveTestSuite struct {
		ServicesTestSuite
	}
)

func (s *RemoveTestSuite) TestRetrievesErrorWhenArgumentNameNotFound() {
	err := Remove(s.ctx)

	s.assert.EqualError(err, ErrServicesRemoveName.Error())
}

func (s *RemoveTestSuite) TestRetrievesErrorWhenApiReturnsError() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{Error: true}}
	s.flag.Parse([]string{"a1"})

	err := Remove(s.ctx)

	s.assert.Error(err)
	s.assert.EqualError(err, ErrServicesRemoveStatus.Error())
}

func (s *RemoveTestSuite) TestRemove() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{}}
	s.flag.Parse([]string{"a1"})

	err := Remove(s.ctx)
	res := s.buf.String()

	s.assert.Nil(err)
	s.assert.Contains(res, "a1")
}

func TestRemoveTestSuite(t *testing.T) {
	suite.Run(t, new(RemoveTestSuite))
}
