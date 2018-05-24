package task

import (
	"fmt"
	"runtime"

	"github.com/emicklei/go-selfdiagnose"
)

// Deprecated: ReportCPU's active goroutines are only calculated during task registering
func ReportCPU() selfdiagnose.ReportMessage {
	cpu := selfdiagnose.ReportMessage{
		Message: fmt.Sprintf("%d CPU available. %d goroutines active", runtime.NumCPU(), runtime.NumGoroutine()),
	}
	cpu.SetComment("Num CPU")
	return cpu
}

type ReportProcessing struct{}

func (r ReportProcessing) Comment() string { return "Processing" }

func (c ReportProcessing) Run(ctx *selfdiagnose.Context, result *selfdiagnose.Result) {
	result.Passed = true
	result.Reason = fmt.Sprintf("%d CPU available. %d goroutines active on %d golang threads", runtime.NumCPU(), runtime.NumGoroutine(), runtime.GOMAXPROCS(-1))
	result.Severity = selfdiagnose.SeverityNone
}
