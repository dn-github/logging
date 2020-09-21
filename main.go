package main

import (
	"github.com/dn-github/logging/logger"
	"gitlab.myteksi.net/gophers/go/commons/util/log/logging"
	"gitlab.myteksi.net/gophers/go/commons/util/log/yall/slog"
	"gitlab.myteksi.net/gophers/go/staples/logging/loggingapi"
)

func main() {
	// util/log/logging
	logConfig := &logging.Config{
		Level: 5,
		Tag:   "action-plan-service",
	}

	logging.Init(logConfig)

	var gkLogger, sLogger loggingapi.Logger
	gkLogger = logging.GetDefaultLogger()
	gkLogger.Warn("gkLogger", "WARNING: %s", "this is a test warning")

	// util/log/yall/slog
	sLogger = logger.New(&slog.Config{
		SyslogTag:       "structuredlog.action-plan-service",
		WorkerCount:     10,
		BufferSize:      10000,
		LogLevel:        3,
		StacktraceLevel: 4,
		Encoding:        slog.CGLSFormat,
		Development:     false,
	})

	sLogger.Warn("sLogger", "WARNING: %s", "this is a test warning")
}
