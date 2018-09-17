package golory

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"sync"
)

var (
	gly *Golory
)

// struct for holding all data
type Golory struct {
	registry sync.Map
	cfg      map[string]interface{}
	booted   bool
}

func init() {
	gly = &Golory{booted: false}
}

// initiate golory components from configuration
// file format support: toml,json,yaml
func Boot(cfg interface{}) error {
	if gly.booted {
		// do clear stuff
		gly.booted = false
	}
	switch cfg.(type) {
	case string:
		if err := bootFromFile(cfg.(string)); err != nil {
			return err
		}
	case []byte:
		if err := bootFromBytes(cfg.([]byte)); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot boot golory configuration, %s", cfg)
	}
	gly.booted = true
	return nil
}

// initate golory components from file
func bootFromFile(path string) error {
	// read file to []byte
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return bootFromBytes(b)
}

// initiate golory components from binary content
func bootFromBytes(b []byte) error {
	if err := parseCfg(b); err != nil {
		return err
	}
	// do initiation
	fmt.Println(gly.cfg)
	// TODO
	// var logc log.Cfg
	// fill(logc, gly.cfg["golory"].(map[interface{}]interface{}))
	// ins := log.Boot(logc)
	// gly.registry.Store("golory.log", ins)
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
