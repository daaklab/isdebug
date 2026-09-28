//go:build debug

package buildtag

import "github.com/daaklab/isdebug"

func init() {
	isdebug.Enabled = true
}
