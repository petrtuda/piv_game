// File: pkg/logger/logger.go

package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger struct {
	*log.Logger
	level LogLevel
}

func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", 0),
		level:  INFO,
	}
}

func (l *Logger) SetOutput(w io.Writer) {
	l.Logger.SetOutput(w)
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) log(level LogLevel, format string, v ...interface{}) {
	if level < l.level {
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	_, file, line, _ := runtime.Caller(2)
	file = file[strings.LastIndex(file, "/")+1:]
	levelStr := [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}[level]

	msg := fmt.Sprintf(format, v...)
	logEntry := fmt.Sprintf("%s [%s] %s:%d: %s", now, levelStr, file, line, msg)

	l.Logger.Output(2, logEntry)

	if level == FATAL {
		os.Exit(1)
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.log(DEBUG, format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.log(INFO, format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.log(WARN, format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.log(ERROR, format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.log(FATAL, format, v...)
}
