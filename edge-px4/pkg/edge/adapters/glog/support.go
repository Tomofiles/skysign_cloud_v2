package glog

import (
	"edge-px4/pkg/edge/domain/common"

	"github.com/golang/glog"
)

// Flush .
func Flush() {
	glog.Flush()
}

// NewSupport .
func NewSupport() common.Support {
	return &supportLog{}
}

type supportLog struct{}

func (s *supportLog) NotifyInfo(format string, args ...interface{}) {
	if len(args) != 0 {
		glog.Infof(format, args...)
	} else {
		glog.Info(format)
	}
}

func (s *supportLog) NotifyError(format string, args ...interface{}) {
	if len(args) != 0 {
		glog.Errorf(format, args...)
	} else {
		glog.Errorf(format)
	}
}
