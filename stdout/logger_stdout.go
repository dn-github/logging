package loggingapi

import (
	"fmt"
	"os"

	"github.com/dn-github/logging/loggingapi"
)

// NewStdOut returns a new logger that will only output log messages to stdout
func NewStdOut() loggingapi.Logger {
	return &stdOutLogger{}
}

// Implement the Logger interface
type stdOutLogger struct{}

// Debug implements Logger interface
func (lr *stdOutLogger) Debug(tag string, message string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "["+tag+"] "+message+"\n", args...)
}

// Info implements Logger interface
func (lr *stdOutLogger) Info(tag string, message string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "["+tag+"] "+message+"\n", args...)
}

// Warn implements Logger interface
func (lr *stdOutLogger) Warn(tag string, message string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "["+tag+"] "+message+"\n", args...)
}

// Error implements Logger interface
func (lr *stdOutLogger) Error(tag string, message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "["+tag+"] "+message+"\n", args...)
}

// Fatal implements Logger interface
func (lr *stdOutLogger) Fatal(tag string, message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "["+tag+"] "+message+"\n", args...)
}

// Security implements Logger interface
func (lr *stdOutLogger) Security(tag string, message string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "["+tag+"] "+message+"\n", args...)
}
