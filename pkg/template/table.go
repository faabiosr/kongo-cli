package template

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	"text/tabwriter"
	tmpl "text/template"
)

type (
	// Table structure for retrieving formatte tables.
	Table struct {
		tmpl   *tmpl.Template
		buf    *bytes.Buffer
		writer *tabwriter.Writer
	}
)

// NewTable retrives an instance of Table template.
func NewTable(format string) (*Table, error) {
	buf := new(bytes.Buffer)
	tmpl, err := tmpl.New("table").Parse(format)

	if err != nil {
		return nil, errors.Wrap(err, ErrTemplateWriting)
	}

	return &Table{
		tmpl,
		buf,
		tabwriter.NewWriter(buf, 20, 1, 3, ' ', 0),
	}, nil
}

// Write writes the template data in the table writer
func (t *Table) Write(data interface{}) error {
	if err := t.tmpl.Execute(t.writer, data); err != nil {
		return errors.Wrap(err, ErrTemplateWriting)
	}

	return nil
}

// Flush writes the buffered data into writer output.
func (t *Table) Flush(writer io.Writer) error {
	if err := t.writer.Flush(); err != nil {
		return errors.Wrap(err, ErrTemplateFlushing)
	}

	if _, err := t.buf.WriteTo(writer); err != nil {
		return errors.Wrap(err, ErrTemplateFlushing)
	}

	return nil
}
