package golory

import (
	"encoding/json"
	"fmt"
	"github.com/1pb-club/golory/log"
	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

var (
	gly *golory
)

// struct for holding all data
type golory struct {
	cfg     *goloryConfig
	handles *handler
	booted  bool
}

type goloryConfig struct {
	Debug bool
	Log   map[string]log.Cfg
}

func init() {
	gly = &golory{
		booted:  false,
		cfg:     &goloryConfig{},
		handles: &handler{},
	}
}

// initiate golory components from configuration
// file format support: toml,json,yaml
// goroutine unsafe
func Boot(cfg interface{}) error {
	if gly.booted {
		// do clear stuff
		gly.booted = false
	}
	switch cfg.(type) {
	case string:
		if err := parseFile(cfg.(string)); err != nil {
			return err
		}
	case []byte:
		if err := parseBytes(cfg.([]byte)); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot boot golory configuration, %s", cfg)
	}
	// do initiation
	gly.init()
	gly.booted = true
	return nil
}

// initate golory components from file
func parseFile(path string) error {
	// read file to []byte
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return parseBytes(b)
}

// initiate golory components from binary content
func parseBytes(b []byte) error {
	if err := parseCfg(b); err != nil {
		return err
	}
	return nil
}

func parseCfg(b []byte) error {
	// try file formats
	var err error
	if err = toml.Unmarshal(b, &gly.cfg); err == nil {
		return nil
	}
	e := wrap(GLY_PARSE_CFG_ERROR, err)
	if err = yaml.Unmarshal(b, &gly.cfg); err == nil {
		return nil
	}
	e = wrap(e, err)
	if err = json.Unmarshal(b, &gly.cfg); err == nil {
		return nil
	}
	return wrap(e, err)
}

func (g *golory) init() {
	g.initLog()
}

// init invoker log
func (g *golory) initLog() {
	// todo generate default log
	if len(g.cfg.Log) > 0 {
		g.handles.log = make(map[string]*log.Logger, 0)
		for key, config := range g.cfg.Log {
			obj := log.Boot(config)
			g.handles.log[key] = obj
			// todo log
		}
	}
}
