package loggingapi

// Logger defines required logging interface
type Logger interface {
	Debug(tag string, message string, args ...interface{})
	Info(tag string, message string, args ...interface{})
	Warn(tag string, message string, args ...interface{})
	Error(tag string, message string, args ...interface{})
	Fatal(tag string, message string, args ...interface{})
	Security(tag string, message string, args ...interface{})
}
