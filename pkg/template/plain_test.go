package template

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	PlainTestSuite struct {
		suite.Suite

		assert *assert.Assertions
	}
)

func (s *PlainTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
}

func (s *PlainTestSuite) TestWriteParsedData() {
	b := &bytes.Buffer{}

	data := struct {
		Type string
	}{"plain"}

	err := NewPlain("{{ .Type }}").Write(b, data)

	s.assert.Contains(b.String(), "plain")
	s.assert.Nil(err)
}

func (s *PlainTestSuite) TestWriteParsedWithInvalidData() {
	b := &bytes.Buffer{}

	data := struct {
		Type string
	}{}

	err := NewPlain("{{ .Type.Test }}").Write(b, data)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrTemplateParsing)
}

func (s *PlainTestSuite) TestWriteParsedWithInvalidContent() {
	b := &bytes.Buffer{}

	data := struct {
		Type string
	}{}

	err := NewPlain(`{{define "foo"}} FOO `).Write(b, data)
	s.assert.Contains(err.Error(), ErrTemplateParsing)

	s.assert.Error(err)
}

func TestPlainTestSuite(t *testing.T) {
	suite.Run(t, new(PlainTestSuite))
}
