package function

import (
	"os"
	"strconv"
)

// Handle a serverless request
func Handle(req []byte) string {
	debug, _ := strconv.ParseBool(os.Getenv("debug"))
	if debug {
		return getName()
	}
	return "OK"
}

func getName() string {
	name, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return name
}
