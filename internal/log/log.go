package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg Config) *zap.Logger {
	return zap.New(
		zapcore.NewCore(getEncoder(cfg), getWriteSyncer(), getLoggerLevel(cfg)),
		getOptions(cfg)...,
	)
}

func getEncoder(cfg Config) zapcore.Encoder { //nolint:ireturn
	var encoderConfig zapcore.EncoderConfig
	if cfg.Production {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	return encoder
}

func getWriteSyncer() zapcore.WriteSyncer { //nolint:ireturn
	return zapcore.Lock(os.Stdout)
}

func getLoggerLevel(cfg Config) zap.AtomicLevel {
	var level zapcore.Level

	if err := level.Set(cfg.Level); err != nil {
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	return zap.NewAtomicLevelAt(level)
}

func getOptions(cfg Config) []zap.Option {
	options := make([]zap.Option, 0)

	if !cfg.Production {
		options = append(options, zap.AddCaller())
		options = append(options, zap.AddStacktrace(zap.ErrorLevel))
	}

	return options
}
