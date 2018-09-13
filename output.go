package golory

import "go.uber.org/zap"

func Log(name string) *zap.Logger {
	// todo default log
	return invokerObj.handler.log[name]
}
