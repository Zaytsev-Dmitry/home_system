package utilities

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TODO создавать логер в зависимости от окружения
func GetLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	return logger
}
