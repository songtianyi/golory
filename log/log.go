package log

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type Cfg struct {
	Debug bool
	Level string
	Path  string
}

// todo for dynamic config & enrich log api
type Logger struct {
	*zap.Logger
}

func Boot(config Cfg) *Logger {
	var js string
	if config.Debug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, config.Level)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, config.Level, config.Path, config.Path)
	}

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	l, err := cfg.Build()
	if err != nil {
		log.Fatal("init logger error: ", err)
	}
	return &Logger{l}

}
