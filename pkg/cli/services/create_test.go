package services

import (
	api "github.com/fabiorphp/kongo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	CreateTestSuite struct {
		ServicesTestSuite
	}
)

func (s *CreateTestSuite) TestRetrievesErrorWhenArgumentNameNotFound() {
	err := Create(s.ctx)

	s.assert.EqualError(err, ErrServicesCreateName.Error())
}

func (s *CreateTestSuite) TestRetrievesErrorWhenArgumentNameIsEmpty() {
	s.flag.Parse([]string{""})

	err := Create(s.ctx)

	s.assert.EqualError(err, ErrServicesCreateName.Error())
}

func (s *CreateTestSuite) TestRetrievesErrorWhenApiReturnsError() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{Error: true}}
	s.flag.Parse([]string{"test"})

	err := Create(s.ctx)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrServicesCreate)
}

func (s *CreateTestSuite) TestCreate() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{}}
	s.flag.Parse([]string{"test"})

	err := Create(s.ctx)
	res := s.buf.String()

	s.assert.Nil(err)
	s.assert.Contains(res, "a1")
}

func (s *CreateTestSuite) TestCreateWithUrl() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{}}
	s.flag.String("url", "ok", "doc")
	s.flag.Parse([]string{"--url=http://getkong.org", "test"})

	err := Create(s.ctx)
	res := s.buf.String()

	s.assert.Nil(err)
	s.assert.Contains(res, "a1")
}

func TestCreateTestSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
