package task

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/emicklei/go-selfdiagnose"
)

type ReportHttpRequest struct{}

func (r ReportHttpRequest) Run(ctx *selfdiagnose.Context, result *selfdiagnose.Result) {
	req, ok := ctx.Variables["http.request"]
	if !ok {
		result.Passed = false
		result.Reason = "missing variable 'http.request'"
		return
	}
	var buf bytes.Buffer
	for k, v := range req.(*http.Request).Header {
		buf.WriteString(fmt.Sprintf("%s = %s<br/>", k, v))
	}
	result.Passed = true
	result.Reason = buf.String()
}

func (r ReportHttpRequest) Comment() string { return "headers of this Http request" }
