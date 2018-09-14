package golory

import "fmt"

// initiate golory components from configuration file or binary content
//
// file format support: toml
func Boot(cfg interface{}) error {
	switch cfg.(type) {
	case string:
		bootFromFile(cfg.(string))
		break
	case []byte:
		bootFromBytes(cfg.([]byte))
		break
	default:
		return fmt.Errorf("cannot boot golory configuration, %s", cfg)
	}
	return nil
}

// initate golory components from confi
func bootFromFile(path string) {
	invokerObj := initInvoker(path)
	invokerObj.initLog()
}

// initiate golory components from configuration binary content
// file format support: toml
func bootFromBytes(b []byte) {

}
