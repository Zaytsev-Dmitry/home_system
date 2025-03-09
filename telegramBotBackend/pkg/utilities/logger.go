package utilities

import "go.uber.org/zap"

func GetLogger() *zap.Logger {
	production, _ := zap.NewProduction()
	return production
}
