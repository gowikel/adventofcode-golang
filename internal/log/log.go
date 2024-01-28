package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/gowikel/adventofcode-golang/internal/conf"
)

type LoggerOption func(*slog.Logger)

// Logger that delegates most calls to slog, but adds some missing
// ones like Fatal
type Logger struct {
	logger *slog.Logger
}

// WithPart adds the part option to the logger
// panics if part is not 1 or 2
func WithPart(part int) LoggerOption {
	if part != 1 && part != 2 {
		panic(
			fmt.Sprintf(
				"WithPart can only be called with 1 or 2, but called with %d instead",
				part,
			),
		)
	}
	return func(l *slog.Logger) {
		*l = *slog.With("part", part)
	}
}

// WithYear adds the year option to the logger
func WithYear(year int) LoggerOption {
	return func(l *slog.Logger) {
		*l = *slog.With("year", year)
	}
}

// WithDay adds the day option to the logger
func WithDay(day int) LoggerOption {
	return func(l *slog.Logger) {
		*l = *slog.With("day", day)
	}
}

// GetLogger returns the default structured logger
// configured with the year and day keys if the
// configuration has been parsed, and the specified
// options.
func GetLogger(options ...LoggerOption) *Logger {
	c := conf.Conf()

	logger := slog.Default()

	if c != nil {
		WithYear(c.Year)(logger)
		WithDay(c.Day)(logger)
	}

	for _, opt := range options {
		opt(logger)
	}

	return &Logger{logger: logger}
}

// Shortcut to slog.Error + os.Exit(1)
func (l *Logger) Fatal(msg string, args ...any) {
	l.logger.Error(msg, args...)
	os.Exit(1)
}

// Delegates to slog.With
func (l *Logger) With(args ...any) *Logger {
	l.logger = l.logger.With(args...)

	return l
}

// Delegates to slog.Debug
func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Delegates to slog.DebugContext
func (l *Logger) DebugContext(
	ctx context.Context,
	msg string,
	args ...any,
) {
	l.logger.DebugContext(ctx, msg, args...)
}

// Delegates to slog.Enabled
func (l *Logger) Enabled(ctx context.Context, level slog.Level) bool {
	return l.logger.Enabled(ctx, level)
}

// Delegates to slog.Error
func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

// Delegates to slog.ErrorContext
func (l *Logger) ErrorContext(
	ctx context.Context,
	msg string,
	args ...any,
) {
	l.logger.ErrorContext(ctx, msg, args...)
}

// Delegates to slog.Handler
func (l *Logger) Handler() slog.Handler {
	return l.logger.Handler()
}

// Delegates to slog.Info
func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Delegates to slog.InfoContext
func (l *Logger) InfoContext(
	ctx context.Context,
	msg string,
	args ...any,
) {
	l.logger.InfoContext(ctx, msg, args...)
}

// Delegates to slog.Log
func (l *Logger) Log(
	ctx context.Context,
	level slog.Level,
	msg string,
	args ...any,
) {
	l.logger.Log(ctx, level, msg, args...)
}

// Delegates to slog.LogAttrs
func (l *Logger) LogAttrs(
	ctx context.Context,
	level slog.Level,
	msg string,
	attrs ...slog.Attr,
) {
	l.logger.LogAttrs(ctx, level, msg, attrs...)
}

// Delegates to slog.Warn
func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Delegates to slog.WarnContext
func (l *Logger) WarnContext(
	ctx context.Context,
	msg string,
	args ...any,
) {
	l.logger.WarnContext(ctx, msg, args...)
}

// Delegates to slog.WithGroup
func (l *Logger) WithGroup(name string) *Logger {
	l.logger = l.logger.WithGroup(name)

	return l
}
