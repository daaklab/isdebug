package env

import (
	"os"

	"github.com/daaklab/isdebug"
)

func init() {
	e := os.Getenv("DEBUG")
	if e != "" {
		isdebug.Enabled = true
	}
}
