package template

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"testing"
)

type (
	PlainTestSuite struct {
		TemplateTestSuite
	}
)

func (s *PlainTestSuite) TestFactoryRetrievesError() {
	_, err := NewPlain(`{{define "foo"}} FOO `)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrTemplateWriting)
}

func (s *PlainTestSuite) TestFactory() {
	tmpl, err := NewPlain(`{{ .Id }}\n`)

	s.assert.Nil(err)
	s.assert.Implements(new(Template), tmpl)
}

func (s *PlainTestSuite) TestWriteRetrievesError() {
	tmpl, _ := NewPlain(`{{ .Id }}\n`)

	data := struct {
		Type string
	}{}

	err := tmpl.Write(data)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrTemplateWriting)
}

func (s *PlainTestSuite) TestFlushFailDuringWriteIntoOutput() {
	tmpl, _ := NewPlain(`{{ .Type }}\n`)

	data := struct {
		Type string
	}{"table"}

	tmpl.Write(data)

	err := tmpl.Flush(new(errorWriter))

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrTemplateFlushing)
}

func (s *PlainTestSuite) TestFlush() {
	tmpl, _ := NewPlain(`{{ .Type }}\n`)

	data := struct {
		Type string
	}{"plain"}

	buf := new(bytes.Buffer)

	s.assert.Nil(tmpl.Write(data))
	s.assert.Nil(tmpl.Flush(buf))
	s.assert.Contains(buf.String(), "plain")
}

func TestPlainTestSuite(t *testing.T) {
	suite.Run(t, new(PlainTestSuite))
}
