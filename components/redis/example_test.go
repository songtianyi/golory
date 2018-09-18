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
