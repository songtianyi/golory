package golory

import (
	"github.com/1pb-club/golory/components/log"
	"github.com/1pb-club/golory/components/redis"
)

// Logger return a log.Logger by name
func Logger(name string) *log.Logger {
	return gly.components.getLogger(name)
}

// Redis return a redis.Client by name
func Redis(name string) *redis.Client {
	return gly.components.getRedis(name)
}
