package task

import (
	"os"

	"github.com/emicklei/go-selfdiagnose"
)

type ReportHostname struct{}

func (r ReportHostname) Comment() string { return "hostname as reported by the os" }

func (r ReportHostname) Run(ctx *selfdiagnose.Context, result *selfdiagnose.Result) {
	h, err := os.Hostname()
	result.Severity = selfdiagnose.SeverityNone
	result.Passed = err == nil
	result.Reason = h
}
