package common

import "github.com/golang/glog"

// Support .
type Support interface {
	NotifyInfo(format string, args ...interface{})
	NotifyError(format string, args ...interface{})
}

// NewSupport .
func NewSupport() Support {
	return &supportLog{}
}

type supportLog struct{}

func (s *supportLog) NotifyInfo(format string, args ...interface{}) {
	if len(args) != 0 {
		glog.Infof(format, args)
	} else {
		glog.Info(format)
	}
}

func (s *supportLog) NotifyError(format string, args ...interface{}) {
	if len(args) != 0 {
		glog.Errorf(format, args)
	} else {
		glog.Errorf(format)
	}
}
