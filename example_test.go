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
   		[golory.logger.default]
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

	fmt.Println(golory.Logger("golory") == nil)
	fmt.Println(golory.Logger("default") == nil)
	// Output:
	// false
	// false
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
	fmt.Println(golory.Redis("default") == nil)
	fmt.Println(golory.Redis("no-this-user") == nil)
	// Output:
	// false
	// true

}

func ExampleMySQL() {
	cfg := `
	[golory]
		[golory.gorm.default]
          debug = false
		  username = "travis"
		  password = ""
		  addr = "127.0.0.1:3306"
		  DBName = "golory"
		  params = {charset="utf8",parseTime="True",loc="Local"}
		  singularTable = true
	`
	if err := golory.Boot([]byte(cfg)); err != nil {
		fmt.Printf("boot golory failed, %s", err)
	}
	fmt.Println(golory.Gorm("default").Err == nil)
	// Output:
	// true
}

func ExamplePostgres() {
	cfg := `
	[golory]
		[golory.gorm.default]
          debug = false
		  username = "travis"
		  password = ""
		  addr = "127.0.0.1:5432"
		  DBName = "golory"
		  singularTable = true
	`
	if err := golory.Boot([]byte(cfg)); err != nil {
		fmt.Printf("boot golory failed, %s", err)
	}
	fmt.Println(golory.Gorm("default").Err == nil)
	// Output:
	// true
}
