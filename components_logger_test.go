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

package golory_test

import (
	"github.com/1pb-club/golory"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestLogger_ParseValidConfig(t *testing.T) {
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
	Convey("解析配置：", t, func() {
		So(
			golory.Boot([]byte(cfg)),
			ShouldBeNil,
		)
	})
}

func TestLogger_LogInfo(t *testing.T) {
	path := "golory.log"
	Convey("写日志：", t, func() {
		logger := golory.LoggerBoot(golory.LoggerCfg{
			Debug: false,
			Level: "INFO",
			Path:  "golory.log",
		})
		logger.Info("something")
		isExist := false
		if _, err := os.Stat(path); err == nil {
			isExist = true
		}
		So(
			isExist,
			ShouldBeTrue,
		)
	})
}
