package golory

import (
	"github.com/1pb-club/golory/log"
)

// Get logger by name
func Logger(name string) *log.Logger {
	return gly.components.getLogger(name)
}
