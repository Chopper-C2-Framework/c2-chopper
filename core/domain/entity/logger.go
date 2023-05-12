package entity

type Logger interface {
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
}
