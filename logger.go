// Copyright 2018 golory Authors @1pb.club. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package golory

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerCfg configurations
type LoggerCfg struct {
	Debug bool
	Level string
	Path  string
}

// LoggerClient wrapped zapper
// TODO: for dynamic config & more log api
type LoggerClient struct {
	*zap.Logger
}

// LoggerBoot logger from config
func (cfg *LoggerCfg) init() (*LoggerClient, error) {
	var js string
	if cfg.Debug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["stdout"],
      "errorOutputPaths": ["stdout"]
      }`, cfg.Level)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "json",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, cfg.Level, cfg.Path, cfg.Path)
	}

	var zcfg zap.Config
	if err := json.Unmarshal([]byte(js), &zcfg); err != nil {
		return nil, fmt.Errorf("init logger error: %s", err)
	}
	zcfg.EncoderConfig = zap.NewProductionEncoderConfig()
	zcfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	l, err := zcfg.Build()
	if err != nil {
		return nil, fmt.Errorf("init logger error: %s", err)
	}
	return &LoggerClient{l}, nil

}

func (s *LoggerClient) close() error {
	// do nothing
	return nil
}
