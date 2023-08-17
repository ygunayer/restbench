package logger

import "github.com/gofiber/fiber/v2/log"

func Trace(v ...interface{}) {
	log.Trace(v...)
}

func Tracef(fmt string, args ...interface{}) {
	log.Tracef(fmt, args...)
}

func Debug(v ...interface{}) {
	log.Debug(v...)
}

func Debugf(fmt string, args ...interface{}) {
	log.Debugf(fmt, args...)
}

func Info(v ...interface{}) {
	log.Info(v...)
}

func Infof(fmt string, args ...interface{}) {
	log.Infof(fmt, args...)
}

func Warn(v ...interface{}) {
	log.Warn(v...)
}

func Warnf(fmt string, args ...interface{}) {
	log.Warnf(fmt, args...)
}

func Error(v ...interface{}) {
	log.Error(v...)
}

func Errorf(fmt string, args ...interface{}) {
	log.Errorf(fmt, args...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func Fatalf(fmt string, args ...interface{}) {
	log.Fatalf(fmt, args...)
}
