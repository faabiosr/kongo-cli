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

	s.assert.Equal("plain", b.String())
	s.assert.Nil(err)
}

func TestPlainTestSuite(t *testing.T) {
	suite.Run(t, new(PlainTestSuite))
}
