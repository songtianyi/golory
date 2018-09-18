package log_test

import (
	"fmt"
	"github.com/1pb-club/golory/log"
)

func ExampleBoot() {
	logger := log.Boot(log.CommonCfg{
		Debug: false,
		Level: "INFO",
		Path:  "golory.log",
	})
	logger.Info("test debug log")
	if logger != nil {
		fmt.Println("OK")
	}

	// Output:
	// OK
}
