package services

import (
	api "github.com/fabiorphp/kongo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	UpdateTestSuite struct {
		ServicesTestSuite
	}
)

func (s *UpdateTestSuite) TestRetrievesErrorWhenArgumentNameNotFound() {
	err := Update(s.ctx)

	s.assert.EqualError(err, ErrServicesUpdateID.Error())
}

func (s *UpdateTestSuite) TestRetrievesErrorWhenArgumentNameIsEmpty() {
	s.flag.Parse([]string{""})

	err := Update(s.ctx)

	s.assert.EqualError(err, ErrServicesUpdateID.Error())
}

func (s *UpdateTestSuite) TestRetrievesErrorWhenApiGetReturnsError() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{GetError: true}}
	s.flag.Parse([]string{"a1"})

	err := Update(s.ctx)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrServicesUpdate)
}

func (s *UpdateTestSuite) TestRetrievesErrorWhenApiUpdateReturnsError() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{Error: true}}
	s.flag.Parse([]string{"test"})

	err := Update(s.ctx)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrServicesUpdate)
}

func (s *UpdateTestSuite) TestUpdateWithAllFlags() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{}}
	s.flag.Int64("connect-timeout", 0, "doc")
	s.flag.String("host", "", "doc")
	s.flag.String("name", "", "doc")
	s.flag.String("path", "", "doc")
	s.flag.Int("port", 0, "doc")
	s.flag.String("protocol", "", "doc")
	s.flag.Int("read-timeout", 0, "doc")
	s.flag.Int("retries", 0, "doc")
	s.flag.Int("write-timeout", 0, "doc")
	s.flag.Parse([]string{
		"--connect-timeout=3000",
		"--host=api.io",
		"--name=api",
		"--path=/v2",
		"--port=9071",
		"--protocol=https",
		"--read-timeout=3000",
		"--retries=7",
		"--write-timeout=3000",
		"test",
	})

	err := Update(s.ctx)
	res := s.buf.String()

	s.assert.Nil(err)
	s.assert.Contains(res, "a1")
}

func (s *UpdateTestSuite) TestUpdateWithUrl() {
	s.ctx.App.Metadata["client"] = &api.Kongo{Services: &MockServices{}}
	s.flag.String("url", "ok", "doc")
	s.flag.Parse([]string{"--url=http://getkong.org", "test"})

	err := Update(s.ctx)
	res := s.buf.String()

	s.assert.Nil(err)
	s.assert.Contains(res, "a1")
}

func TestUpdateTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
