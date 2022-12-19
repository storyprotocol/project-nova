package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsage(t *testing.T) {
	assert.Equal(t, "log levels: info|debug|warn|error", Usage)
}
