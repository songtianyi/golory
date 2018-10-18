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

func TestMySQL_ParseConfig(t *testing.T) {
	cfg := `
	[golory]
		[golory.mysql.default]
          debug = false
		  username = "travis"
		  password = ""
		  addr = "127.0.0.1:3306"
		  DBName = "golory"
		  params = {charset="utf8",parseTime="True",loc="Local"}
		  singularTable = true
	`
	Convey("Parse MySQL cfg:", t, func() {
		So(
			Boot([]byte(cfg)),
			ShouldBeNil,
		)
	})
}

func TestMySQL_InitConn(t *testing.T) {
	cfg1 := MySQLCfg{
		Addr:     "127.0.0.1:3306",
		Username: "travis",
		Password: "",
		DBName:   "golory",
	}
	Convey("Init mysql conn", t, func() {
		So(
			cfg1.init().ConnectionErr,
			ShouldBeNil,
		)
	})
}
