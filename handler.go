package golory

import (
	"github.com/1pb-club/golory/components/log"
)

type handler struct {
	loggers map[string]*log.Logger
}

func newHandler() *handler {
	return &handler{
		loggers: make(map[string]*log.Logger),
	}
}
func (s *handler) setLogger(k string, v *log.Logger) {
	s.loggers[k] = v
}

func (s *handler) getLogger(k string) *log.Logger {
	l := s.loggers[k]
	return l
}
