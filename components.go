package golory

import "github.com/1pb-club/golory/log"

// exported functions for component usages
func Logger(name string) *log.Logger {
	// todo default log
	return gly.handles.log[name]
}
