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
    		path = "./default.log"
	`
	if err := golory.Boot([]byte(cfg)); err != nil {
		fmt.Printf("boot golory failed, %s", err)
	}
	fmt.Println(golory.Logger("asdf"))
	// Output:
	// <nil>
}
