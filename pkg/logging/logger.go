package logging

import (
	"log/slog"
	"os"

	"github.com/artemKapitonov/avito_test_task/pkg/logging/slogpretty"
)

const (
	defaultLevel      = LevelInfo
	defaultAddSource  = true
	defaultIsJSON     = true
	defaultSetDefault = true
)

// LoggerOptions is type for cofiguration logging.Logger.
type LoggerOptions struct {
	IsLocal    bool   `yaml:"is_local"` // Enable slogpretty for coloful and readable logs.
	Level      string `yaml:"level"`
	AddSource  bool   `yaml:"add_source"`
	IsJSON     bool   `yaml:"is_json"`
	SetDefault bool   `yaml:"set_default"`
}

// NewLogger creats new logging.Logger who implement slog.Logger.
func NewLogger(cfg LoggerOptions) *Logger {
	if cfg.IsLocal {
		return setupPrettySlog()
	}

	var level = setLoggerLevel(cfg.Level)

	handlerOpt := setupHandlerOptions(level, cfg.AddSource)
	var handler Handler
	if cfg.IsJSON {
		handler = NewJSONHandler(os.Stdout, handlerOpt)
	} else {
		handler = NewTextHandler(os.Stdout, handlerOpt)
	}

	logger := New(handler)

	if cfg.SetDefault {
		slog.SetDefault(logger)
	}

	return logger
}

// Default creates new default slog.Logger.
func Default() *Logger {
	return slog.Default()
}

func setupPrettySlog() *Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func setLoggerLevel(lvl string) Level {
	var level Level
	switch lvl {
	case "debug":
		level = -4
	case "info":
		level = 0
	case "warn":
		level = 4
	case "error":
		level = 8
	}

	return level
}

func setupHandlerOptions(level Level, AddSource bool) *HandlerOptions {
	return &HandlerOptions{AddSource: AddSource, Level: level}
}
