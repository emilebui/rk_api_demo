package logger

import "testing"

func TestInitFakeLogger(t *testing.T) {
	Logger = InitFakeLogger()
	Logger.Info("Testing Worked!!!")
}
