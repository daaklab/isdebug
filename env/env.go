package env

import (
	"os"

	"github.com/daakdev/isdebug"
)

func init() {
	e := os.Getenv("DEBUG")
	if e != "" {
		isdebug.Enabled = true
	}
}
