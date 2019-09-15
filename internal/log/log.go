package log

import "fmt"

type nilLogger struct {}
func (nl *nilLogger) Printf(format string, v ...interface{}) {}

var errorLog Logger = new(nilLogger)
var debugLog Logger = new(nilLogger)

type Logger interface {
	Printf(format string, v ...interface{})
}

func SetErrorLogger(logger Logger) { errorLog = logger }
func SetDebugLogger(logger Logger) { debugLog = logger }

func Debug(format string, v ...interface{}) {
	debugLog.Printf(format, v...)
}
func Error(format string, v ...interface{}) {
	errorLog.Printf(fmt.Errorf(format, v...).Error())
}