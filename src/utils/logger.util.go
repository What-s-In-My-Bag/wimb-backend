package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func GetLogger() *zap.Logger {
	if Logger == nil {
		Logger, _ = zap.NewProduction()
		defer Logger.Sync()
		return Logger
	}
	return Logger
}
