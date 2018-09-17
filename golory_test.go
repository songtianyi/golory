package golory

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBoot(t *testing.T) {
	// TODO
	cfg := `
	[golory]
    	[golory.log]
    		debug = true
    		level = "info"
    		path = "./default.log"
	`
	err := Boot([]byte(cfg))
	assert.Equal(t, err, nil)
}
