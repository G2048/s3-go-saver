package configs

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

// type Logslevel struct {
// 	debug slog.Level
// 	info  slog.Level
// 	warn  slog.Level
// 	error slog.Level
// }

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
	// buildInfo, _ := debug.ReadBuildInfo()
	var Logger = slog.New(slog.NewJSONHandler(os.Stderr, &options)).With(
		slog.String("app", "yandex_storage"),
		slog.String("version", "1.0.0"),
		// slog.Group("program_info",
		// slog.Int("pid", os.Getpid()),
		// slog.String("go_version", buildInfo.GoVersion),
		// ),
	)
	slog.SetDefault(Logger)
	return Logger
}
