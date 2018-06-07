package template

import (
	"bytes"
	"github.com/stretchr/testify/suite"
	"testing"
	"text/tabwriter"
	tmpl "text/template"
)

type (
	TableTestSuite struct {
		TemplateTestSuite
	}
)

func (s *TableTestSuite) TestFactoryRetrievesError() {
	_, err := NewTable(`{{define "foo"}} FOO `)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrTemplateWriting)
}

func (s *TableTestSuite) TestFactory() {
	tmpl, err := NewTable(`{{ .Id }}\n`)

	s.assert.Nil(err)
	s.assert.Implements(new(Template), tmpl)
}

func (s *TableTestSuite) TestWriteRetrievesError() {
	tmpl, _ := NewTable(`{{ .Id }}\n`)

	data := struct {
		Type string
	}{}

	err := tmpl.Write(data)

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrTemplateWriting)
}

func (s *TableTestSuite) TestFlushFailDuringWriteTable() {
	tmpl := &Table{
		tmpl.New("testing"),
		new(bytes.Buffer),
		tabwriter.NewWriter(new(errorWriter), 20, 1, 3, ' ', 0),
	}

	err := tmpl.Flush(new(bytes.Buffer))

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrTemplateFlushing)
}

func (s *TableTestSuite) TestFlushFailDuringWriteIntoOutput() {
	tmpl, _ := NewTable(`{{ .Type }}\n`)

	data := struct {
		Type string
	}{"table"}

	tmpl.Write(data)

	err := tmpl.Flush(new(errorWriter))

	s.assert.Error(err)
	s.assert.Contains(err.Error(), ErrTemplateFlushing)
}

func (s *TableTestSuite) TestFlush() {
	tmpl, _ := NewTable(`{{ .Type }}\n`)

	data := struct {
		Type string
	}{"table"}

	buf := new(bytes.Buffer)

	s.assert.Nil(tmpl.Write(data))
	s.assert.Nil(tmpl.Flush(buf))
	s.assert.Contains(buf.String(), "table")
}

func TestTableTestSuite(t *testing.T) {
	suite.Run(t, new(TableTestSuite))
}
