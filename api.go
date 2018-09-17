package golory

import "github.com/1pb-club/golory/log"

func Log(name string) *log.Logger {
	// todo default log
	return gly.handler.log[name]
}