package template

import (
	"io"
)

var (
	// ErrTemplateFlushing retrieves an error message when template flushing fails.
	ErrTemplateFlushing = "Template flushing error"

	// ErrTemplateWriting retrieves an error message when template writing fails.
	ErrTemplateWriting = "Template writing error"
)

type (
	// Template is responsible for formatting CLI output
	Template interface {
		// Write writes the data.
		Write(data interface{}) error

		// Flush writes the data into writer output.
		Flush(writer io.Writer) error
	}
)
