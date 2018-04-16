package template

import (
	"io"
)

type (
	// Template is responsible for formatting CLI ouput
	Template interface {
		// Write writes the template data in the IO
		Write(writer io.Writer, data interface{}) error
	}
)
