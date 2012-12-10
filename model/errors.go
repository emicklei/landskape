package model

import "github.com/emicklei/hopwatch"
func Debug(args ...interface{}) {
	hopwatch.Display(args...).Break()
}

const (
	MISMATCH_ID    = 4001 // first special BadRequest(400)
	MISMATCH_SCOPE = 4002 // first special BadRequest(400)
)
