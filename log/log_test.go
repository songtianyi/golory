package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := Boot(Cfg{
		Debug: false,
		Level: "INFO",
		Path:  "golory.log",
	})
	logger.Info("test debug log")
}
