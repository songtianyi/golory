package golory

import (
	"github.com/1pb-club/golory/log"
)

// Get logger by name
func Logger(name string) *log.Logger {
	// todo default log
	return gly.components.getLogger(name)
}
