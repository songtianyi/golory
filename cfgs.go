package golory

import "github.com/1pb-club/golory/log"

type invokerConfig struct {
	Golory goloryConfig
}
type goloryConfig struct {
	Debug bool
	Log   map[string]log.Cfg
}
