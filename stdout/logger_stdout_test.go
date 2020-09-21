package stdout

import (
	"testing"

	"github.com/dn-github/logging/loggingapi"

	"github.com/stretchr/testify/assert"
)

func TestStdOutLogger_implements(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*loggingapi.Logger)(nil), NewStdOut())
}

func TestStdOutLogger_Methods(t *testing.T) {
	t.Parallel()

	testLoggerMethods(NewStdOut())
}
