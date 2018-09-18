package golory

import (
	"github.com/1pb-club/golory/components/log"
)

// Logger return a log.Logger by name
func Logger(name string) *log.Logger {
	return gly.components.getLogger(name)
}

// Redis do nothing
func Redis() {

}
