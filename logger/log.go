/*
Package the logger wraps any logger to simplify changing logger or inject any other packages as needed.
*/
package logger

import (
	"context"
	"os"

	"log/slog"
)

var DefaultLogger = func() *Logger {
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(handler)
	return &Logger{logger: logger}
}

// Logger represents an active logging object. Used as wrapper of slog.
type Logger struct {
	logger *slog.Logger
}

// New initializes a new logger.
func New(isProd bool) *Logger {
	var handler slog.Handler
	if isProd {
		handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	logger := slog.New(handler)
	return &Logger{logger: logger}
}

// AppendFields adds structured fields to the logger.
func (l *Logger) AppendFields(fields map[string]interface{}) {
	for k, v := range fields {
		l.logger = l.logger.With(k, v)
	}
}

// DEFAULT LOGGER METHODS

func Info(msg string) {
	DefaultLogger().logger.Info(msg)
}

func Debug(msg string) {
	DefaultLogger().logger.Debug(msg)
}

func Fatal(err error, msg string) {
	DefaultLogger().logger.Error(msg, "error", err)
	os.Exit(1)
}

func Error(err error, msg string) {
	DefaultLogger().logger.Error(msg, "error", err)
}

func Infof(fmt string, v ...interface{}) {
	DefaultLogger().logger.InfoContext(context.Background(), fmt, "args", v)
}

func Debugf(fmt string, v ...interface{}) {
	DefaultLogger().logger.DebugContext(context.Background(), fmt, "args", v)
}

func Errorf(err error, fmt string, v ...interface{}) {
	DefaultLogger().logger.ErrorContext(context.Background(), fmt, "error", err, "args", v)
}

// INSTANCE METHODS

func (l *Logger) Info(msg string) {
	if l == nil {
		return
	}
	l.logger.Info(msg)
}

func (l *Logger) Debug(msg string) {
	if l == nil {
		return
	}
	l.logger.Debug(msg)
}

func (l *Logger) Error(err error, msg string) {
	if l == nil {
		return
	}
	l.logger.Error(msg, "error", err)
}

func (l *Logger) Fatal(err error, msg string) {
	if l == nil {
		return
	}
	l.logger.Error(msg, "error", err)
	os.Exit(1)
}

func (l *Logger) Infof(fmt string, v ...interface{}) {
	if l == nil {
		return
	}
	l.logger.InfoContext(context.Background(), fmt, "args", v)
}

func (l *Logger) Debugf(fmt string, v ...interface{}) {
	if l == nil {
		return
	}
	l.logger.DebugContext(context.Background(), fmt, "args", v)
}

func (l *Logger) Errorf(err error, fmt string, v ...interface{}) {
	if l == nil {
		return
	}
	l.logger.ErrorContext(context.Background(), fmt, "error", err, "args", v)
}

func (l *Logger) InfofWithFields(fields map[string]interface{}, fmt string, v ...interface{}) {
	if l == nil {
		return
	}
	log := l.logger
	for k, v := range fields {
		log = log.With(k, v)
	}
	log.InfoContext(context.Background(), fmt, "args", v)
}

func (l *Logger) SendWithFields(fields map[string]interface{}) {
	if l == nil {
		return
	}
	log := l.logger
	for k, v := range fields {
		log = log.With(k, v)
	}
	log.InfoContext(context.Background(), "Message sent with fields")
}

func (l *Logger) SendErrorWithFields(err error, fields map[string]interface{}) {
	if l == nil {
		return
	}
	log := l.logger
	for k, v := range fields {
		log = log.With(k, v)
	}
	log.ErrorContext(context.Background(), "Error sent with fields", "error", err)
}

// Print prints interface message to logger writer with Info level.
func (l *Logger) Print(v ...interface{}) {
	if l == nil {
		return
	}
	l.logger.InfoContext(context.Background(), "%v", "args", v)
}
