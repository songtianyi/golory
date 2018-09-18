package golory

import "github.com/1pb-club/golory/log"

type invokerHandler struct {
	log map[string]*log.Logger
}

func (g *golory) init() {
	
}

// init invoker log
func (g *golory) initLog() {
	// todo generate default log
	if len(g.cfg.Golory.Log) > 0 {
		g.handler.log = make(map[string]*log.Logger, 0)
		for key, config := range g.cfg.Golory.Log {
			obj := log.Boot(config)
			g.handler.log[key] = obj
			// todo log
		}
	}
}