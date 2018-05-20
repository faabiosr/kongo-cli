package template

import (
	"bytes"
	"github.com/pkg/errors"
	"io"
	tmpl "text/template"
)

var (
	// ErrTemplateParsing retrieves an error message when template parsing fails.
	ErrTemplateParsing = "Template parsing error"
)

type (
	// Plain structure for retrieving simple templates
	Plain struct {
		content string
		tmpl    *tmpl.Template
		buf     *bytes.Buffer
	}
)

// NewPlain retrieves an instanceo of Plain template
func NewPlain(content string) *Plain {
	return &Plain{
		content,
		tmpl.New("plain"),
		bytes.NewBufferString(""),
	}
}

// Write writes the template data in the IO
func (p *Plain) Write(writer io.Writer, data interface{}) error {
	tm, err := p.tmpl.Parse(p.content)

	if err != nil {
		return errors.Errorf("%s: %v", ErrTemplateParsing, err)
	}

	err = tm.Execute(p.buf, data)

	if err != nil {
		return errors.Errorf("%s: %v", ErrTemplateParsing, err)
	}

	p.buf.WriteString("\n\n")

	_, err = p.buf.WriteTo(writer)

	return err
}
