package logger

import (
	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

const (
	// Info level
	Info = "info"
	// Debug level
	Debug = "debug"
	// Trace level - Shows the same logs as debug level but
	// with detailed stack trace
	Trace = "trace"
)

var configMap = map[string]zap.Config{
	Info: {
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    baseEncoderConfig(),
	},
	Debug: {
		Encoding:      "json",
		OutputPaths:   []string{"stderr"},
		Level:         zap.NewAtomicLevelAt(zapcore.DebugLevel),
		EncoderConfig: baseEncoderConfig(),
	},
	Trace: {
		Encoding:      "json",
		OutputPaths:   []string{"stderr"},
		Level:         zap.NewAtomicLevelAt(zapcore.DebugLevel),
		EncoderConfig: traceEncoderConfig(),
	},
}

// New constructs a zap logger
func New(logLevel string) (*zap.Logger, error) {
	loggerConfig, ok := configMap[logLevel]
	if !ok {
		loggerConfig = configMap[Info]
	}

	return loggerConfig.Build()
}

func baseEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     zapcore.OmitKey,
		FunctionKey:   zapcore.OmitKey,
		MessageKey:    "msg",
		StacktraceKey: zapcore.OmitKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.TimeEncoderOfLayout(
			"2006-01-02T15:04:05Z",
		),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func traceEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   "function",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.TimeEncoderOfLayout(
			"2006-01-02T15:04:05Z",
		),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
