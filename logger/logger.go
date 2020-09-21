package logger

import (
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/dn-github/logging/loggingapi"
	"gitlab.myteksi.net/gophers/go/commons/util/log/yall/slog"
	"gitlab.myteksi.net/gophers/go/commons/util/tags"
)

const (
	fnKey           = "fn"
	defaultStatsKey = "service"
)

type logger struct {
	slogger *slog.Logger
}

var initOnce sync.Once

// New initializes logger from ucm config
func New(config *slog.Config) loggingapi.Logger {
	if config == nil {
		slogger, _, _ := slog.Stdout()
		return &logger{slogger}
	}
	slogger, _, _ := slog.FromConfig(config)

	initOnce.Do(func() {
		Default = &logger{slogger}
	})

	return &logger{slogger}
}

// Error error method of a logger, can be used when you want to use logger explicitly
func (l *logger) Error(eventType, message string, args ...tags.Tag) {
	args = append(args, tags.CustomTag(fnKey, caller(2)))
	l.slogger.Error(eventType, message, args...)
}

// Warn warn method of a logger, can be used when you want to use logger explicitly
func (l *logger) Warn(eventType, message string, args ...tags.Tag) {
	args = append(args, tags.CustomTag(fnKey, caller(2)))
	l.slogger.Warn(eventType, message, args...)
}

// Info info method of a logger, can be used when you want to use logger explicitly
func (l *logger) Info(eventType, message string, args ...tags.Tag) {
	l.slogger.Info(eventType, message, args...)
}

// Debug debug method of a logger, can be used when you want to use logger explicitly
func (l *logger) Debug(eventType, message string, args ...tags.Tag) {
	l.slogger.Debug(eventType, message, args...)
}

// With with method of a logger, can be used when you want to use logger explicitly
func (l *logger) With(args ...tags.Tag) *slog.Logger {
	return l.slogger.With(args...)
}

// Caller formats the caller
func caller(skip int) string {
	_, fn, line, _ := runtime.Caller(skip)
	idx := strings.LastIndex(fn, "gophers/go/")
	if idx > 0 {
		idx += 11
	}

	return fn[idx:] + "." + strconv.FormatInt(int64(line), 10)
}
