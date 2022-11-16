package logger

import (
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"go.uber.org/zap"
)

var Logger *rkentry.LoggerEntry

// InitFakeLogger initializes the fake logger for testing purposes
func InitFakeLogger() *rkentry.LoggerEntry {

	logger, _ := zap.NewProduction()

	return &rkentry.LoggerEntry{
		Logger: logger,
	}
}
