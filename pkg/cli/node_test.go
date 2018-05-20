package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeCommand(t *testing.T) {
	assert.Len(t, Node().Subcommands, 1)
}
