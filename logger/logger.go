package logger

//import (
//	"fmt"
//	"io"
//	"os"
//	"runtime"
//	"sync"
//	"time"
//)
//
//const (
//	green   = "\033[97;42m"
//	white   = "\033[90;47m"
//	yellow  = "\033[90;43m"
//	red     = "\033[97;41m"
//	blue    = "\033[97;44m"
//	magenta = "\033[97;45m"
//	cyan    = "\033[97;46m"
//	reset   = "\033[0m"
//)
//
//type logLevel uint8
//
//const (
//	debug logLevel  = iota
//	info
//	warning
//	err
//	fatal
//)
//
//type marsLog struct {
//	mux sync.Mutex
//	level logLevel
//	out    io.Writer  // destination for output
//	buf []byte
//}
//
//func init() {
//	log = &marsLog{level:debug,mux:sync.Mutex{},out:os.Stderr}
//}
//func SetDebug() {
//
//}
//
//var log *marsLog
//func Debugf(format string, args ...interface{}) {
//	if log.level <= debug {
//		log.outPut("[D]",format,args...)
//	}
//}
////
////func Infof(format string, args ...interface{}) {
////	if log.level <= info {
////		log.outPut("[I]",format,args)
////	}
////}
////
////func Warningf(format string, args ...interface{}) {
////	if log.level <= warning {
////		log.outPut("[W]",format,args)
////	}
////}
////
////func Errorf(format string, args ...interface{}) {
////	if log.level <= err {
////		log.outPut("[E]",format,args)
////	}
////}
////
////func Fatalf(format string, args ...interface{}) {
////	if log.level <= fatal {
////		log.outPut("[F]",format,args)
////		os.Exit(1)
////	}
////}
////func Debug(args ...interface{}) {
////	if log.level <= debug {
////		log.outPut("[D]",args...)
////	}
////}
////
//func Info(args ...interface{}) {
//	if log.level <= info {
//		log.outPut(args...)
//	}
//}
////
////func Warning(args ...interface{}) {
////	if log.level <= warning {
////		log.outPut("[W]",...args)
////	}
////}
////
////func Error(args ...interface{}) {
////	if log.level <= err {
////		log.outPut("[E]",args...)
////	}
////}
////
////func Fatal(args ...interface{}) {
////	if log.level <= fatal {
////		log.outPut("[F]",args...)
////		os.Exit(1)
////	}
////}
//func (l *marsLog)outPut(s string) error{
//	now := time.Now() // get this early.
//	var file string
//	var line int
//
//	var ok bool
//	_, file, line, ok = runtime.Caller(2)
//	if !ok {
//		file = "???"
//		line = 0
//	}
//	l.mux.Lock()
//	defer l.mux.Unlock()
//	l.buf = append(l.buf, s...)
//	if len(s) == 0 || s[len(s)-1] != '\n' {
//		l.buf = append(l.buf, '\n')
//	}
//	_, err := l.out.Write(l.buf)
//	return err
//}