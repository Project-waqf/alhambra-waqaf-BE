package helper

import "go.uber.org/zap"

func Logger() zap.Logger {
	cfg := zap.NewProductionConfig()

	cfg.DisableStacktrace = true

	logger, _ := cfg.Build()
	return *logger
}