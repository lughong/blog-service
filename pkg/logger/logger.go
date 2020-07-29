package logger

import (
	"fmt"
	"log"

	"github.com/natefinch/lumberjack"

	"github.com/lughong/blog-service/global"
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
	level     string
	files     Files
	callers   []string
}

func NewLogger() *Logger {
	l := log.New(
		&lumberjack.Logger{
			Filename: fmt.Sprintf(
				"%s/%s.%s",
				global.AppSetting.LogSavePath,
				global.AppSetting.LogFileName,
				global.AppSetting.LogFileExt,
			),
			MaxSize:    500,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		},
		"",
		0,
	)

	return &Logger{
		newLogger: l,
		files:     make(Files),
		callers:   make([]string, 0),
	}
}

func (l *Logger) clone() *Logger {
	c := *l
	return &c
}

func (l *Logger) WithLevel(level string) *Logger {
	cl := l.clone()
	cl.level = level
	return cl
}

func (l *Logger) WithFiles(files Files) *Logger {
	cl := l.clone()
	for k, f := range files {
		cl.files[k] = f
	}

	return cl
}

func (l *Logger) WithCaller() {}

func (l *Logger) JsonFormat()       {}
func (l *Logger) Output(msg string) {}

func (l *Logger) Debug(msg string) {
	l.WithLevel(LevelDebug.String()).Output(msg)
}

func (l *Logger) Debugf(format string, msg ...interface{}) {
	l.WithLevel(LevelDebug.String()).Output(fmt.Sprintf(format, msg...))
}

func (l *Logger) Info(msg string) {
	l.WithLevel(LevelInfo.String()).Output(msg)
}

func (l *Logger) Infof(format string, msg ...interface{}) {
	l.WithLevel(LevelInfo.String()).Output(fmt.Sprintf(format, msg...))
}

func (l *Logger) Warn(msg string) {
	l.WithLevel(LevelWarn.String()).Output(msg)
}

func (l *Logger) Warnf(format string, msg ...interface{}) {
	l.WithLevel(LevelWarn.String()).Output(fmt.Sprintf(format, msg...))
}

func (l *Logger) Error(msg string) {
	l.WithLevel(LevelError.String()).Output(msg)
}

func (l *Logger) Errorf(format string, msg ...interface{}) {
	l.WithLevel(LevelError.String()).Output(fmt.Sprintf(format, msg...))
}

func (l *Logger) Fatal(msg string) {
	l.WithLevel(LevelFatal.String()).Output(msg)
}

func (l *Logger) Fatalf(format string, msg ...interface{}) {
	l.WithLevel(LevelFatal.String()).Output(fmt.Sprintf(format, msg...))
}

func (l *Logger) Panic(msg string) {
	l.WithLevel(LevelPanic.String()).Output(msg)
}

func (l *Logger) Panicf(format string, msg ...interface{}) {
	l.WithLevel(LevelPanic.String()).Output(fmt.Sprintf(format, msg...))
}
