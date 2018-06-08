package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServicesCommand(t *testing.T) {
	assert.Len(t, Services().Subcommands, 4)
}
