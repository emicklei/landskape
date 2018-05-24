package selfdiagnose

import (
	"testing"
	"time"
)

type long struct {
	BasicTask
}

func (l long) Run(ctx *Context, result *Result) {
	time.Sleep(10 * time.Second)
	result.Passed = true
	result.Reason = "waited out"
}
func TestThatLongTasksAreTimedout(t *testing.T) {
	l := new(long)
	l.SetTimeout(5 * time.Second)
	reg := &Registry{}
	reg.Register(l)
	rr := new(recordingReporter)
	reg.Run(rr)
	if len(rr.results) == 0 {
		t.Fatal("no results")
	}
	if got, want := rr.results[0].Passed, false; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
