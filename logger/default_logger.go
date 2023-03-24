package logger

import (
	"log"
)

var (
	defaultLog = &Logger{}
)

func init() {
	if err := defaultLoggerInit(); err != nil {
		log.Printf("Default log initialization failed, err: %v", err)
	}
}

func defaultLoggerInit() error {
	if err := defaultLog.Init(LevelALL, "", "", ""); err != nil {
		return err
	}
	defaultLog.SetCallDepth(4)
	return nil
}

func SetDefaultLogOutPrefix(outPrefix string) {
	defaultLog.SetOutPrefix(outPrefix)
}

func SetDefaultLogLevel(logLevelDesc string) {
	lv, exist := FindLogLevelByDesc(logLevelDesc)
	if !exist {
		lv = LevelALL
	}
	defaultLog.SetLogLevel(lv)
}

func All(v ...interface{}) {
	defaultLog.All("", v...)
}

func AllFmt(format string, v ...interface{}) {
	defaultLog.All(format, v...)
}

func Debug(v ...interface{}) {
	defaultLog.Debug("", v...)
}

func DebugFmt(format string, v ...interface{}) {
	defaultLog.Debug(format, v...)
}

func Info(v ...interface{}) {
	defaultLog.Info("", v...)
}

func InfoFmt(format string, v ...interface{}) {
	defaultLog.Info(format, v...)
}

func Warning(v ...interface{}) {
	defaultLog.Warning("", v...)
}

func WarnFmt(format string, v ...interface{}) {
	defaultLog.Warning(format, v...)
}

func Error(v ...interface{}) {
	defaultLog.Error("", v...)
}

func ErrorFmt(format string, v ...interface{}) {
	defaultLog.Error(format, v...)
}
