package task

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/emicklei/go-selfdiagnose"
)

// go test -v -test.run=TestCheckDirectory
func TestCheckDirectory(t *testing.T) {
	loc := filepath.Join(os.TempDir(), "selfdiagnose_test")
	testfile, err := os.Create(loc)
	if err != nil {
		t.Fatal("unable to create testfile:" + err.Error())
	}
	testfile.Close()

	{
		reg := &Registry{}
		reg.Register(CheckDirectory{Path: os.TempDir(), CanAppend: true})
		rr := new(recordingReporter)
		reg.Run(rr)
		if len(rr.results) == 0 {
			t.Fatal("no results")
		}
		if got, want := rr.results[0].Passed, true; got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	{
		reg := &Registry{}
		check := CheckDirectory{Path: filepath.Join(os.TempDir(), "missing"), CanAppend: true}
		if got, want := check.Severity(), SeverityCritical; got != want {
			t.Errorf("got %v want %v", got, want)
		}
		reg.Register(check)
		rr := new(recordingReporter)
		reg.Run(rr)
		if len(rr.results) == 0 {
			t.Fatal("no results")
		}
		if got, want := rr.results[0].Passed, false; got != want {
			t.Errorf("got %v want %v", got, want)
		}
		if got, want := rr.results[0].Severity, SeverityCritical; got != want {
			t.Errorf("got <%v> want <%v>", got, want)
		}
	}
}

func ExampleCheckDirectory() {
	check := CheckDirectory{Path: "/tmp", CanAppend: true}
	check.SetComment("something")
	Register(check)
}
