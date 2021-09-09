package common

// Support .
type Support interface {
	NotifyInfo(format string, args ...interface{})
	NotifyError(format string, args ...interface{})
}
