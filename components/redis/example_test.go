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

package redis_test

import (
	"fmt"
	"github.com/1pb-club/golory/components/redis"
)

func ExampleBoot() {
	c := redis.Boot(redis.CommonCfg{
		Addr: "127.0.0.1:6379",
	})
	n, err := c.Exists("4774EF3B-F3B8-4D17-AB18-AC39AB7B4329-s").Result()
	fmt.Printf("%d:%v", n, err)
	// Output:
	// 0:<nil>
}
