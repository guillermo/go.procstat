package procstat

import (
	"os"
	"testing"
)

func TestUpdate(t *testing.T) {

	s := &Stat{Pid: os.Getpid()}
	err := s.Update()
	if err != nil {
		t.Error(err)
	}

}
