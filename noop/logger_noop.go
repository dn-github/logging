package noop

import "github.com/dn-github/logging/loggingapi"

// NewNoop returns a new No Op Logger.
func NewNoop() loggingapi.Logger {
	return &noOpLogger{}
}

// Implement the Logger interface and "no op" all calls to the logger.
// This is used when no other logger is specified
type noOpLogger struct{}

// Debug implements Logger interface
func (lr *noOpLogger) Debug(_ string, _ string, _ ...interface{}) {
	// do nothing
}

// Info implements Logger interface
func (lr *noOpLogger) Info(_ string, _ string, _ ...interface{}) {
	// do nothing
}

// Warn implements Logger interface
func (lr *noOpLogger) Warn(_ string, _ string, _ ...interface{}) {
	// do nothing
}

// Error implements Logger interface
func (lr *noOpLogger) Error(_ string, _ string, _ ...interface{}) {
	// do nothing
}

// Fatal implements Logger interface
func (lr *noOpLogger) Fatal(_ string, _ string, _ ...interface{}) {
	// do nothing
}

// Security implements Logger interface
func (lr *noOpLogger) Security(_ string, _ string, _ ...interface{}) {
	// do nothing
}
