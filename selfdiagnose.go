package main

import (
	"os"

	selfdiagnose "github.com/emicklei/go-selfdiagnose"
	"github.com/emicklei/go-selfdiagnose/task"
)

func initSelfdiagnose() {
	// add http handlers for /internal/selfdiagnose.(html|json|xml)
	selfdiagnose.AddInternalHandlers()
	selfdiagnose.Register(task.ReportHttpRequest{})
	selfdiagnose.Register(task.ReportHostname{})
	selfdiagnose.Register(task.ReportCPU())

	selfdiagnose.Register(task.ReportVariables{map[string]interface{}{
		"GOOGLE_CLOUD_PROJECT": os.Getenv("GOOGLE_CLOUD_PROJECT")}})
}
