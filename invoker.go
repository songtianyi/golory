package golory

import (
	"encoding/json"
	"fmt"
	"github.com/1pb-club/golory/log"
	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
)

var invokerObj *invoker

type invoker struct {
	config  *invokerConfig
	handler *invokerHandler
}

type invokerHandler struct {
	log map[string]*zap.Logger
}

type invokerConfig struct {
	Golory goloryConfig
}

type goloryConfig struct {
	Debug bool
	Log   map[string]log.LogStruct
}

func initInvoker(configFile string) *invoker {
	invokerObj = &invoker{
		config:  &invokerConfig{},
		handler: &invokerHandler{},
	}
	err := invokerObj.parseConfig(configFile)
	if err != nil {
		panic(err.Error())
	}
	invokerObj.debug(invokerObj.config)
	return invokerObj
}

func (i *invoker) parseConfig(configFile string) (err error) {
	// Set defaults.
	_, err = os.Stat(configFile)
	if err != nil {
		return
	}

	var content []byte
	content, err = ioutil.ReadFile(configFile)
	if err != nil {
		return
	}

	fileExt := filepath.Ext(configFile)
	switch fileExt {
	case ".toml":
		err = toml.Unmarshal(content, i.config)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(content, i.config)
	case ".json":
		err = json.Unmarshal(content, i.config)
	default:
		err = fmt.Errorf("file ext %v unsupported", fileExt)
	}
	return
}

func (i *invoker) initLog() {
	if len(i.config.Golory.Log) > 0 {
		i.handler.log = make(map[string]*zap.Logger, 0)
		for key, config := range i.config.Golory.Log {
			obj := log.InitLog(config)
			i.handler.log[key] = obj
			i.debug("initLog name: " + key)
		}
	}
}

func (i *invoker) debug(name interface{}) {
	if invokerObj.config.Golory.Debug {
		fmt.Println(name)
	}
}
