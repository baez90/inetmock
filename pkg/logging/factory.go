package logging

import (
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggingConfig = zap.NewProductionConfig()
)

func ConfigureLogging(
	level zap.AtomicLevel,
	developmentLogging bool,
	initialFields map[string]interface{},
) {
	loggingConfig.Level = level
	loggingConfig.Development = developmentLogging
	if initialFields != nil {
		loggingConfig.InitialFields = initialFields
	}
}

func ParseLevel(levelString string) zap.AtomicLevel {
	switch strings.ToLower(levelString) {
	case "debug":
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zapcore.FatalLevel)
	default:
		return zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
}

func CreateLogger() (Logger, error) {
	if zapLogger, err := loggingConfig.Build(zap.AddCallerSkip(2)); err != nil {
		return nil, err
	} else {
		return NewLogger(zapLogger), nil
	}
}

func CreateTestLogger(tb testing.TB) Logger {
	return &testLogger{
		tb:      tb,
		encoder: zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
	}
}
