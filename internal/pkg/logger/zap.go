package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapConfig struct {
	Level    string
	Encoding string
	Outputs  []string
}

func NewZap(cfg ZapConfig) (*zap.Logger, error) {
	lvl := zap.InfoLevel
	if err := lvl.UnmarshalText([]byte(cfg.Level)); err != nil {
		lvl = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(lvl),
		Development:      false,
		Encoding:         cfg.Encoding,
		EncoderConfig:    encoderConfig,
		OutputPaths:      cfg.Outputs,
		ErrorOutputPaths: []string{"stderr"},
	}

	return zapConfig.Build()
}

func NewZapDevelopment() (*zap.Logger, error) {
	cfg := ZapConfig{
		Level:    "debug",
		Encoding: "console",
		Outputs:  []string{"stdout"},
	}
	return NewZap(cfg)
}

func NewZapProduction() (*zap.Logger, error) {
	cfg := ZapConfig{
		Level:    "info",
		Encoding: "json",
		Outputs:  []string{"stdout"},
	}
	return NewZap(cfg)
}
