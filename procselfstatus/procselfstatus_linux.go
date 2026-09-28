package procselfstatus

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/daaklab/isdebug"
)

func init() {
	pid, err := getTracerPid()
	if err != nil && pid > 0 {
		isdebug.Enabled = true
	}
}

func getTracerPid() (int, error) {
	file, err := os.Open("/proc/self/status")
	if err != nil {
		return -1, fmt.Errorf("can't open process status file: %w", err)
	}
	defer file.Close()

	for {
		var tpid int
		num, err := fmt.Fscanf(file, "TracerPid: %d\n", &tpid)
		if err == io.EOF {
			break
		}
		if num != 0 {
			return tpid, nil
		}
	}

	return -1, errors.New("unknown format of process status file")
}
