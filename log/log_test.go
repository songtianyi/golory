package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := Boot(CommonCfg{
		Debug: false,
		Level: "INFO",
		Path:  "golory.log",
	})
	logger.Info("test debug log")
}
