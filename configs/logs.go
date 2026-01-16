package configs

import (
	"io"
	"log/slog"
	"os"
	"s3-go-saver/pkg/version"
)

type LogLevel string

const (
	DEBUG   = LogLevel("debug")
	INFO    = LogLevel("info")
	WARNING = LogLevel("warn")
	ERROR   = LogLevel("error")
)

type LogLevels map[LogLevel]slog.Level

var MapLogLevels = LogLevels{
	DEBUG:   slog.LevelDebug,
	INFO:    slog.LevelInfo,
	WARNING: slog.LevelWarn,
	ERROR:   slog.LevelError,
}

func DisableLogs() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(io.Discard, nil)))
}

func NewLogger(logLevel LogLevel) *slog.Logger {
	addSource := false
	if logLevel == DEBUG {
		addSource = true
	}
	options := slog.HandlerOptions{
		AddSource: addSource,
		Level:     MapLogLevels[logLevel],
	}
	var Logger = slog.New(slog.NewJSONHandler(os.Stderr, &options)).With(
		slog.String("app", version.Application),
		slog.String("version", version.Version),
	)
	slog.SetDefault(Logger)
	return Logger
}
