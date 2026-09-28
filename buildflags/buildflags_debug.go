//go:build debug

package buildtag

import "github.com/daakdev/isdebug"

func init() {
	isdebug.Enabled = true
}
