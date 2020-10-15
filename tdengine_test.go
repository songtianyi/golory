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
	"testing"
)

func TestTDengine_ParseConfig(t *testing.T) {
	cfg := `
	[golory]
		[golory.tdengine.default]
		  Addr = "127.0.0.1:6030"
		  Username = "root"
	          Password = "taosdata"
		[golory.tdengine.remote]
		  Addr = "192.168.1.99:6030"
	`
	Convey("Parse tdengine cfg:", t, func() {
		So(
			Boot([]byte(cfg)),
			ShouldBeNil,
		)
	})
}

func TestTDengine_InitConn(t *testing.T) {
	cfg1 := TDengineCfg{
		Addr:     "127.0.0.1:6380",
		Username: "root",
		Password: "tasodata",
		DBName:   "top1",
	}
	Convey("Init tdengine conn", t, func() {
		c, err := cfg1.init()
		So(
			err,
			ShouldBeNil,
		)
		c.close()
	})
}
