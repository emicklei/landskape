// +build !windows

package task

import (
	"fmt"
	"syscall"

	"github.com/emicklei/go-selfdiagnose"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type ReportDiskusage struct {
	Path string
}

func (r ReportDiskusage) Run(ctx *selfdiagnose.Context, result *selfdiagnose.Result) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(r.Path, &fs)
	if err != nil {
		result.Passed = false
		result.Reason = err.Error()
		return
	}

	result.Passed = true
	result.Reason = fmt.Sprintf("Size: %.2f GB<br />Free: %.2f GB", float64(fs.Blocks*uint64(fs.Bsize))/float64(GB), float64(fs.Bfree*uint64(fs.Bsize))/float64(GB))
	return
}

func (r ReportDiskusage) Comment() string { return fmt.Sprintf("disk usage of %q", r.Path) }
