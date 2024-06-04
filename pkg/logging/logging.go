package logging

import (
	"io"
	"log/slog"
	"os"
)

// Logger struct for writing logs with slog.Logger in io.Writer.
type Logger struct {
	Logger *slog.Logger
	Writer io.Writer
}

const (
	writerFile   = "file"
	writerStdout = "stdout"
	levelLocal   = "local"
	levelDev     = "dev"
	levelProd    = "prod"
	handlerJSON  = "json"
	handlerText  = "text"
)

type Config struct {
	Level   string `yaml:"level" env:"LOG_LEVEL"`
	Handler string `yaml:"handler" env:"LOG_HANDLER"`
	Writer  string `yaml:"writer" env:"LOG_WRITER"`
}

// New setup new logging.Logger with config params.
func New(cfg Config) *Logger {

	lvl := selectLoggerLevel(cfg.Level)

	writer := selectLoggerWriter(cfg.Writer)

	logger := selectLoggerHandler(loggerOptions{lvl: lvl, w: writer, h: cfg.Handler})

	return &Logger{
		Logger: logger,
		Writer: writer,
	}
}

func selectLoggerWriter(w string) io.Writer {

	var writer io.Writer

	switch w {

	case writerFile:
		writer = getLogFile()

	case writerStdout:
		writer = os.Stdout

	default:
		writer = os.Stdout
	}

	return writer
}

func selectLoggerLevel(level string) slog.Level {
	var lvl slog.Level

	switch level {
	case levelLocal:
		lvl = slog.LevelDebug

	case levelDev:
		lvl = slog.LevelInfo

	case levelProd:
		lvl = slog.LevelError

	default:
		lvl = slog.LevelDebug
	}

	return lvl
}

type loggerOptions struct {
	h   string
	lvl slog.Level
	w   io.Writer
}

func selectLoggerHandler(opts loggerOptions) *slog.Logger {
	var logger *slog.Logger

	switch opts.h {
	case handlerJSON:
		logger = slog.New(slog.NewJSONHandler(opts.w, &slog.HandlerOptions{Level: opts.lvl}))

	case handlerText:
		logger = slog.New(slog.NewTextHandler(opts.w, &slog.HandlerOptions{Level: opts.lvl}))

	default:
		logger = slog.New(slog.NewTextHandler(opts.w, &slog.HandlerOptions{Level: opts.lvl}))
	}

	return logger
}

func getLogFile() *os.File {
	err := os.RemoveAll("logs")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("logs", 0777)
	if err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	return logFile
}
