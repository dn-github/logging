package noop

import (
	"testing"

	"github.com/dn-github/logging/loggingapi"

	"github.com/stretchr/testify/assert"
)

func TestNoOpLogger_implements(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*loggingapi.Logger)(nil), &noOpLogger{})
}

func TestNoOpLogger_Methods(t *testing.T) {
	t.Parallel()

	testLoggerMethods(NewNoop())
}
