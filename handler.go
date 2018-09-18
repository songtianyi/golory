package golory

import (
	"github.com/1pb-club/golory/log"
	"sync"
)

type handler struct {
	loggers map[string]*log.Logger
	mu      sync.RWMutex
}

func newHandler() *handler {
	return &handler{
		loggers: make(map[string]*log.Logger),
	}
}
func (s *handler) setLogger(k string, v *log.Logger) {
	s.mu.Lock()
	s.loggers[k] = v
	s.mu.Unlock()
}

func (s *handler) getLogger(k string) *log.Logger {
	s.mu.RLock()
	l := s.loggers[k]
	s.mu.RUnlock()
	return l
}
