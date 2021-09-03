package common

import "github.com/golang/glog"

// Support .
type Support interface {
	NotifyInfo(format string, args ...interface{})
	NotifyError(format string, args ...interface{})
}

// NewSupport .
func NewSupport() Support {
	return &support{}
}

type support struct{}

func (s *support) NotifyInfo(format string, args ...interface{}) {
	glog.Infof(format, args)
}

func (s *support) NotifyError(format string, args ...interface{}) {
	glog.Errorf(format, args)
}
