package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level uint8
type Files map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return "Unknown level"
	}
}

type Logger struct {
	newLogger *log.Logger
	level     Level
	files     Files
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	return &Logger{
		newLogger: log.New(w, prefix, flag),
	}
}

func (l *Logger) clone() *Logger {
	cl := *l

	return &cl
}

func (l *Logger) WithLevel(lvl Level) *Logger {
	cl := l.clone()

	cl.level = lvl
	return cl
}

func (l *Logger) WithFiles(files Files) *Logger {
	cl := l.clone()
	if cl.files == nil {
		cl.files = make(Files)
	}

	for k, f := range files {
		cl.files[k] = f
	}

	return cl
}

func (l *Logger) WithCaller(skip int) *Logger {
	cl := l.clone()

	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		cl.callers = []string{
			fmt.Sprintf("%s, %d, %s", file, line, f.Name()),
		}
	}

	return cl
}

func (l *Logger) WithCallerFrames() {

}

func (l *Logger) JsonFormat(message string) Files {
	data := make(Files)
	data["level"] = l.level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers

	if len(l.files) > 0 {
		for k, f := range l.files {
			data[k] = f
		}
	}

	return data
}

func (l *Logger) Output(message string) {
	body, _ := json.Marshal(l.JsonFormat(message))
	content := string(body)

	switch l.level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.WithLevel(LevelWarn).Output(fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.WithLevel(LevelWarn).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.WithLevel(LevelError).Output(fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.WithLevel(LevelError).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.WithLevel(LevelPanic).Output(fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.WithLevel(LevelPanic).Output(fmt.Sprintf(format, v...))
}
