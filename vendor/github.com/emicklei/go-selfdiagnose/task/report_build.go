package task

import (
	"fmt"

	"github.com/emicklei/go-selfdiagnose"
)

// Reports Build and Date
type ReportBuildAndDate struct {
	Name    string
	Version string
	Date    string
}

func (r ReportBuildAndDate) Run(ctx *selfdiagnose.Context, result *selfdiagnose.Result) {
	result.Passed = true
	result.Reason = fmt.Sprintf("%s - version:%s date:%s", r.Name, r.Version, r.Date)
}

func (r ReportBuildAndDate) Comment() string {
	return "Build information"
}
