package golory

import (
	"fmt"
)

// Exported errors
var (
	// Error occurred when parse configuration
	GLY_PARSE_CFG_ERROR = fmt.Errorf("parse cfg failed")
)

// Join strings
func wrap(e error, cause error) error {
	return fmt.Errorf("%s, %s", e.Error(), cause.Error())
}
