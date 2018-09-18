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
