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

func TestRedis_ParseConfig(t *testing.T) {
	cfg := `
	[golory]
		[golory.redis.default]
		  Addr = "127.0.0.1:6379"
		[golory.redis.user]
		  Addr = "127.0.0.1:6479"
	`
	Convey("Parse redis cfg:", t, func() {
		So(
			Boot([]byte(cfg)),
			ShouldBeNil,
		)
	})
}

func TestRedis_InitConn(t *testing.T) {
	cfg1 := RedisCfg{
		Addr: "127.0.0.1:6379",
	}
	cfg2 := RedisCfg{
		Addr: "127.0.0.1:6380",
	}
	Convey("Init redis conn", t, func() {
		c, _ := cfg1.init()
		So(c.Ping().Err(),
			ShouldBeNil,
		)
		c.close()
		c, _ = cfg2.init()
		So(c.Ping().Err(),
			ShouldNotBeNil,
		)
		c.close()
	})
}
