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
	ll := *l

	return &ll
}

func (l *Logger) WithLevel(lvl Level) *Logger {
	ll := l.clone()

	ll.level = lvl
	return ll
}

func (l *Logger) WithFiles(files Files) *Logger {
	ll := l.clone()
	if ll.files == nil {
		ll.files = make(Files)
	}

	for k, f := range files {
		ll.files[k] = f
	}

	return ll
}

func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()

	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{
			fmt.Sprintf("%s, %d, %s", file, line, f.Name()),
		}
	}

	return ll
}

func (l *Logger) WithCallerFrames() *Logger {
	minDepth := 0
	maxDepth := 25

	pc := make([]uintptr, maxDepth)
	n := runtime.Callers(minDepth, pc)
	frames := runtime.CallersFrames(pc[:n])

	callers := make([]string, 0)
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s, %d, %s", frame.File, frame.Line, frame.Function))
	}

	ll := l.clone()
	ll.callers = callers

	return ll
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
