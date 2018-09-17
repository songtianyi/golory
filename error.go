package golory

import (
	"fmt"
)

// exported errors

var (
	GLY_PARSE_CFG_ERROR = fmt.Errorf("parse cfg failed")
)

// join strings
func wrap(e error, cause error) error {
	return fmt.Errorf("%s, %s", e.Error(), cause.Error())
}
