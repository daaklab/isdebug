package cmdargs

import (
	"os"

	"github.com/daaklab/isdebug"
)

func init() {
	for _, arg := range os.Args {
		if arg == "--debug" {
			isdebug.Enabled = true
		}
	}
}
