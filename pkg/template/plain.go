package template

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	tmpl "text/template"
)

type (
	// Plain structure for retrieving simple templates
	Plain struct {
		tmpl *tmpl.Template
		buf  *bytes.Buffer
	}
)

// NewPlain retrieves an instanceo of Plain template
func NewPlain(content string) (*Plain, error) {
	tmpl, err := tmpl.New("plain").Parse(content)

	if err != nil {
		return nil, errors.Wrap(err, ErrTemplateWriting)
	}

	return &Plain{
		tmpl,
		bytes.NewBufferString(""),
	}, nil
}

// Write writes the template data in the buffer
func (p *Plain) Write(data interface{}) error {
	if err := p.tmpl.Execute(p.buf, data); err != nil {
		return errors.Wrap(err, ErrTemplateWriting)
	}

	return nil
}

// Flush writes the buffered data into writer output.
func (p *Plain) Flush(writer io.Writer) error {
	if _, err := p.buf.WriteTo(writer); err != nil {
		return errors.Wrap(err, ErrTemplateFlushing)
	}

	return nil
}
