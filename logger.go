package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevel()
	if err := config.Level.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}
	return config.Build()
}

func Filter(logger *zap.Logger, level string) *zap.Logger {
	config := logger.Core().Config()
	config.Level = zap.NewAtomicLevel()
	if err := config.Level.UnmarshalText([]byte(level)); err != nil {
		return logger
	}
	return logger.WithOptions(zap.AddFilter(func(ent zapcore.Entry) bool {
		return ent.Level >= config.Level.Level()
	}))
}