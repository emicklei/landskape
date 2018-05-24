package selfdiagnose

import "testing"

func TestReportMessage(t *testing.T) {
	check := ReportMessage{Message: "Hello World"}

	reg := &Registry{}
	reg.Register(check)
	rr := new(recordingReporter)
	reg.Run(rr)
	if len(rr.results) == 0 {
		t.Fatal("no results")
	}
	if got, want := rr.results[0].Passed, true; got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
