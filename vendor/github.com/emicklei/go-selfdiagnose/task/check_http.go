package task

// Copyright 2013,2016 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	. "github.com/emicklei/go-selfdiagnose"
)

// CheckHTTP send a http.Request and check the expected status code.
type CheckHTTP struct {
	BasicTask
	Method       string
	URL          string
	StatusCode   int
	ShowResponse bool
	HTTPClient   *http.Client
}

// NewCheckHTTP returns a CheckHTTP for threadsafe use of a Request.
func NewCheckHTTP(method string, urlstr string, expectedStatus int) *CheckHTTP {
	return &CheckHTTP{
		Method:     method,
		URL:        urlstr,
		StatusCode: expectedStatus,
	}
}

// Run sends the request and updates the result.
func (c *CheckHTTP) Run(ctx *Context, result *Result) {
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	request, err := http.NewRequest(c.Method, c.URL, nil)
	if err != nil {
		result.Passed = false
		result.Reason = fmt.Sprintf("%s %s => %s", c.Method, c.URL, err.Error())
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		result.Passed = false
		result.Reason = fmt.Sprintf("%s %s => %s", c.Method, c.URL, err.Error())
		return
	}
	defer resp.Body.Close()
	result.Passed = resp.StatusCode == c.StatusCode
	summary := fmt.Sprintf("%s %s => %s", request.Method, request.URL.String(), resp.Status)
	if c.ShowResponse {
		var buf bytes.Buffer
		buf.WriteString(summary)
		buf.WriteString("\n\n")
		io.Copy(&buf, resp.Body)
		summary = buf.String()
	}
	result.Reason = summary
}
