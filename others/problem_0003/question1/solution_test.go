package question1

import (
	"fmt"
	"testing"
	"time"
)

// normal case
type Workload struct {
	timeout time.Duration
}

func (w *Workload) Process() {
	time.Sleep(w.timeout)
	fmt.Print("Process exit normally\n")
}

func TestProcess(t *testing.T) {
	workload := &Workload{timeout: 1 * time.Second}
	err := Question1(workload)
	if err != nil {
		t.Error(err)
	}
}

func TestTimeout(t *testing.T) {
	workload := &Workload{timeout: 6 * time.Second}
	err := Question1(workload)
	if err == nil {
		t.Errorf("should be timeout\n")
	}
}
