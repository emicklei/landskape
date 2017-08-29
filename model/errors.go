package model

import (
	"log"
)

func Debug(args ...interface{}) {
	log.Println(args...)
}
func Display(args ...interface{}) {
	log.Println(args...)
}

const (
	MISMATCH_ID    = 4001 // first special BadRequest(400)
	MISMATCH_SCOPE = 4002 // first special BadRequest(400)
)
