package golory

import (
	"github.com/1pb-club/golory/log"
)

// Logger return a log.Logger by name
func Logger(name string) *log.Logger {
	return gly.components.getLogger(name)
}
