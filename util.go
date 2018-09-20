package golory

import "go.uber.org/zap"

// debugLog golory debug is true,then show the msg
func debugLog(module string,msg string)  {
	if gly.cfg.Golory.Debug {
		glyLogger.Info(msg,zap.String("module",module))
	}
}
