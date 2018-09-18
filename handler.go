package golory

import (
	"github.com/1pb-club/golory/components/log"
	"github.com/1pb-club/golory/components/redis"
)

type handler struct {
	loggers      map[string]*log.Logger
	redisClients map[string]*redis.Client
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

func (s *handler) setRedis(k string, v *redis.Client) {
	s.redisClients[k] = v
}

func (s *handler) getRedis(k string) *redis.Client {
	c := s.redisClients[k]
	return c
}
