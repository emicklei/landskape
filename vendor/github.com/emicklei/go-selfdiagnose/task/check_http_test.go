package task

import (
	"net/http"
	"testing"

	. "github.com/emicklei/go-selfdiagnose"
)

func TestCheckHTTP(t *testing.T) {
	check := NewCheckHTTP("GET", "http://ernestmicklei.com", http.StatusOK)
	check.SetComment("blog access")

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

func ExampleCheckHTTP() {
	check := NewCheckHTTP("GET", "http://ernestmicklei.com", http.StatusOK)
	check.SetComment("blog access")
	Register(check)
}
