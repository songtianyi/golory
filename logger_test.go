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
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"os"
	"testing"
)

func TestLogger_ParseConfig(t *testing.T) {
	cfg := `
		[golory]
		debug = true
    	[golory.logger.golory]
    		debug = true
    		level = "error"
    		path = "golory.log"
		[golory.logger.default]
    		debug = true
    		level = "info"
    		path = "default.log"
	`
	Convey("Parse logger cfg:", t, func() {
		So(
			Boot([]byte(cfg)),
			ShouldBeNil,
		)
	})
}

func TestLogger_WriteLog(t *testing.T) {
	path := "TestLogger_WriteLog.log"
	Convey("Write log", t, func() {

		cfg := LoggerCfg{
			Debug: false,
			Level: "INFO",
			Path:  "TestLogger_WriteLog.log",
		}
		logger, err := cfg.init()
		So(
			err,
			ShouldBeNil,
		)
		logger.Info("something")
		isExist := false
		if _, err := os.Stat(path); err == nil {
			isExist = true
			// remove it
			if err := os.Remove(path); err != nil {
				log.Printf("remove %s failed, %s\n", path, err)
			}
		}
		So(
			isExist,
			ShouldBeTrue,
		)
	})
}
