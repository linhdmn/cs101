package logging

import (
	"context"
)

var (
	l Logger = NullLogger{}
)

// SetDefaultLogger set the default logger
func SetDefaultLogger(logger Logger) {
	if logger == nil {
		logger = NullLogger{}
	}
	l = logger
}

type Logger interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})

	WithFields(fields map[string]interface{}) Logger
	WithField(k string, v interface{}) Logger
}

type loggerKey struct{}

// IntoContext return a new context with the logger injected
func IntoContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// FromContext return the logger from a context if any,
// if no logger in the context, it returns a default Logger
func FromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(loggerKey{}).(Logger); ok {
		return l
	}

	return l
}

// WithFields inject new logger with given key-value pairs and return new context,
// this is a shortcut for get a Logger FromContext, add field and inject it back to the context
func WithFields(ctx context.Context, fields map[string]interface{}) (context.Context, Logger) {
	logger := FromContext(ctx).WithFields(fields)
	return IntoContext(ctx, logger), logger
}

// WithField inject new logger with given key-value pair and return new context,
// this is a shortcut for get a Logger FromContext, add field and inject it back to the context
func WithField(ctx context.Context, k string, v interface{}) (context.Context, Logger) {
	logger := FromContext(ctx).WithField(k, v)
	return IntoContext(ctx, logger), logger
}

func Copy(dst context.Context, src context.Context) context.Context {
	return IntoContext(dst, FromContext(src))
}

// NullLogger implement Logger interface but does nothing, useful for testing
type NullLogger struct{}

// Info implement Logger
func (n NullLogger) Info(_ ...interface{}) {}

// Infof implement Logger
func (n NullLogger) Infof(_ string, _ ...interface{}) {}

// Warn implement Logger
func (n NullLogger) Warn(_ ...interface{}) {}

// Warnf implement Logger
func (n NullLogger) Warnf(_ string, _ ...interface{}) {}

// Error implement Logger
func (n NullLogger) Error(_ ...interface{}) {}

// Errorf implement Logger
func (n NullLogger) Errorf(_ string, _ ...interface{}) {}

// Fatal implement Logger
func (n NullLogger) Fatal(_ ...interface{}) {}

// Fatalf implement Logger
func (n NullLogger) Fatalf(_ string, _ ...interface{}) {}

// WithField implement Logger
func (n NullLogger) WithField(_ string, _ interface{}) Logger { return n }

// WithFields implement Logger
func (n NullLogger) WithFields(_ map[string]interface{}) Logger { return n }
