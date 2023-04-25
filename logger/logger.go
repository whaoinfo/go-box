package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type LogLevel uint8

const (
	DefaultCallDepth = 3
)

const (
	LevelALL LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

var (
	levelDescMap = map[LogLevel]string{
		LevelALL: "[ALL]", LevelDebug: "[DEBUG]", LevelInfo: "[INFO]",
		LevelWarning: "[WARNING]", LevelError: "[ERROR]", LevelFatal: "[FATAL]",
	}
)

type Logger struct {
	logger    *log.Logger
	level     LogLevel
	callDepth int
	dirPath   string
	fileName  string
	logFile   *os.File
	isStdOut  bool
	outPrefix string
}

func (t *Logger) Init(level LogLevel, outPrefix, dirPath, fileName string) error {
	t.level = level
	t.callDepth = DefaultCallDepth

	if err := t.bindLogFile(dirPath, fileName); err != nil {
		return err
	}

	t.outPrefix = outPrefix
	t.logger = log.New(t.logFile, t.outPrefix, log.Llongfile|log.Ldate|log.Ltime)
	return nil
}

func (t *Logger) SetCallDepth(depth int) {
	t.callDepth = depth
}

func (t *Logger) SetLogLevel(level LogLevel) {
	t.level = level
}

func (t *Logger) SetOutPrefix(outPrefix string) {
	t.outPrefix = outPrefix
	if t.logger == nil {
		return
	}
	t.logger.SetPrefix(outPrefix)
}

func (t *Logger) GetLogLevel() LogLevel {
	return t.level
}

func (t *Logger) bindLogFile(dirPath, fileName string) error {
	if dirPath == "" || fileName == "" {
		t.isStdOut = true
		t.logFile = os.Stdout
		return nil
	}

	if !t.checkPathExist(dirPath) {
		return fmt.Errorf("%s path dose not exist", dirPath)
	}

	filePath := path.Join(dirPath, fmt.Sprintf("%s.log", fileName))
	f, openErr := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if openErr != nil {
		return openErr
	}

	t.logFile = f
	return nil
}

func (t *Logger) checkPathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func (t *Logger) Release() error {
	if t.isStdOut || t.logFile == nil {
		return nil
	}

	return t.logFile.Close()
}

func (t *Logger) Output(level LogLevel, format string, v ...interface{}) error {
	if level < t.level {
		return nil
	}

	if t.logger == nil {
		return nil
	}

	var content string
	if format != "" {
		content = fmt.Sprintf(format, v...)
	} else {
		content = fmt.Sprintln(v...)
	}

	return t.logger.Output(t.callDepth, strings.Join([]string{levelDescMap[level], content}, " "))
}

func (t *Logger) All(format string, v ...interface{}) {
	_ = t.Output(LevelALL, format, v...)
}

func (t *Logger) Debug(format string, v ...interface{}) {
	_ = t.Output(LevelDebug, format, v...)
}

func (t *Logger) Info(format string, v ...interface{}) {
	_ = t.Output(LevelInfo, format, v...)
}

func (t *Logger) Warning(format string, v ...interface{}) {
	_ = t.Output(LevelWarning, format, v...)
}

func (t *Logger) Error(format string, v ...interface{}) {
	_ = t.Output(LevelError, format, v...)
}

func FindLogLevelByDesc(levelDesc string) (LogLevel, bool) {
	levelDesc = fmt.Sprintf("[%s]", levelDesc)
	for lv, desc := range levelDescMap {
		if levelDesc == desc {
			return lv, true
		}
	}

	return 0, false
}
