package template

import (
	"io"
	tmpl "text/template"
)

type (
	// Plain structure for retrieving simple templates
	Plain struct {
		content string
		tmpl    *tmpl.Template
	}
)

// NewPlain retrieves an instanceo of Plain template
func NewPlain(content string) *Plain {
	return &Plain{
		content,
		tmpl.New("plain"),
	}
}

// Write writes the template data in the IO
func (p *Plain) Write(writer io.Writer, data interface{}) error {
	tm, err := p.tmpl.Parse(p.content)

	if err != nil {
		return err
	}

	return tm.Execute(writer, data)
}
