package golory_test

import (
	"fmt"
	"github.com/1pb-club/golory"
)

func ExampleBoot() {
	cfg := `
	[golory]
    	[golory.log]
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
    	[golory.log]
    		debug = true
    		level = "info"
    		path = "default.log"
	`
	if err := golory.Boot([]byte(cfg)); err != nil {
		fmt.Printf("boot golory failed, %s", err)
	}
	// TODO
	fmt.Println(golory.Logger("asdf"))
	// Output:
	// <nil>
}

func ExampleRedis() {
	cfg := `
	[golory]
		[redis]
		  Addr = "127.0.0.1:6379"
	`
	if err := golory.Boot([]byte(cfg)); err != nil {
		fmt.Printf("boot golory failed, %s", err)
	}
	// TODO
	fmt.Println(golory.Redis("sdf"))
	// Output:
	// <nil>

}
