package helper

import "go.uber.org/zap"

func Logger() zap.Logger {
	logger, _ := zap.NewProduction()
	return *logger
}