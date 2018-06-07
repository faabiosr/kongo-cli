package template

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type (
	TemplateTestSuite struct {
		suite.Suite

		assert *assert.Assertions
	}

	errorWriter struct{}
)

func (e *errorWriter) Write([]byte) (int, error) {
	return 0, errors.New("cannot write")
}

func (s *TemplateTestSuite) SetupTest() {
	s.assert = assert.New(s.T())
}
