package configs

import (
	"io"
	"log/slog"
	"os"
	"s3-go-saver/pkg/version"
	"strings"
)

var LogLevels = [4]string{"debug", "info", "warn", "error"}
var logLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func DisableLogs() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(io.Discard, nil)))
}

func NewLogger(level string) *slog.Logger {
	level = strings.ToLower(level)
	addSource := false
	if level == "debug" {
		addSource = true
	}
	options := slog.HandlerOptions{
		AddSource: addSource,
		Level:     logLevels[level],
	}
	var Logger = slog.New(slog.NewJSONHandler(os.Stderr, &options)).With(
		slog.String("app", version.Application),
		slog.String("version", version.Version),
	)
	slog.SetDefault(Logger)
	return Logger
}
