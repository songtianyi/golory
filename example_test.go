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
	"fmt"
	"github.com/1pb-club/golory"
)

func ExampleBoot() {
	cfg := `
	[golory]
    	[golory.log.default]
    		debug = true
    		level = "info"
    		path = "./default.log"
	`
	fmt.Println(golory.Boot([]byte(cfg)))
	// Output:
	// <nil>
}

func ExampleLogger() {
	cfg := `
		[golory]
    	[golory.logger.golory]
    		debug = true
    		level = "error"
    		path = "golory.log"
		[golory.logger.default]
    		debug = true
    		level = "info"
    		path = "default.log"
	`
	if err := golory.Boot([]byte(cfg)); err != nil {
		fmt.Printf("boot golory failed, %s", err)
	}

	// TODO
	fmt.Println(golory.Logger("golory"))
	fmt.Println(golory.Logger("default"))
	// Output:
	// <nil>
}

func ExampleRedis() {
	cfg := `
	[golory]
		[golory.redis.default]
		  Addr = "127.0.0.1:6379"
		[golory.redis.user]
		  Addr = "127.0.0.1:6479"
	`
	if err := golory.Boot([]byte(cfg)); err != nil {
		fmt.Printf("boot golory failed, %s", err)
	}
	// TODO
	fmt.Println(golory.Redis("default"))
	// Output:
	// <nil>

}

func ExampleMySql() {
	cfg := `
	[golory]
		[golory.mysql.default]
		  UserName = "root"
		  PassWord = "root"
		  Addr = "127.0.0.1:3306"
		  Name = "golory"
	`
	// [golory.mysql.user]
	// UserName = "root"
	// PassWord = "root"
	// Addr = "127.0.0.1:3306"
	// Name = "test"
	// TablePrefix = "test_"
	// Dsn = {charset = "utf8",parseTime = "True",loc = "Local"}
	if err := golory.Boot([]byte(cfg)); err != nil {
		fmt.Printf("boot golory failed, %s", err)
	}
	// TODO
	fmt.Println(golory.MySql("default"))
	// Output:
	// <nil>

}
