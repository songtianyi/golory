package redis

import (
	gredis "github.com/go-redis/redis"
)

// CommonCfg wrapped go-redis options and some other golory options
type CommonCfg struct {
	Addr string
	// TODO cluster options
	// TODO logger option?
}

// Client wrapped go-redis RedisClient
type Client struct {
	*gredis.Client
}

// Boot return a *redis.Client
func Boot(cfg CommonCfg) *Client {
	return &Client{
		gredis.NewClient(&gredis.Options{
			Addr: cfg.Addr,
		}),
	}
}
