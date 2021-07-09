package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	logLevel  = "LOG_LEVEL"
	logOutput = "LOG_OUTPUT"
)

var (
	log = logger{}
)

type logger struct {
	l *zap.Logger
}

type BookstoreLogger interface {
	Printf(string, ...interface{})
	Print(v ...interface{})
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutput()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	var l *zap.Logger
	if l, err = logConfig.Build(); err != nil {
		panic(err)
	}
	log.l = l
}

func getOutput() string {
	output := strings.TrimSpace(os.Getenv(logOutput))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(logLevel))) {
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "debug":
		return zap.DebugLevel
	default:
		return zap.InfoLevel
	}

}

func Info(msg string, tags ...zap.Field) {
	log.l.Info(msg, tags...)
	log.l.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.Error(err))
	log.l.Error(msg, tags...)
	log.l.Sync()
}

func Fatal(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.Error(err))
	log.l.Fatal(msg, tags...)
	log.l.Sync()
}

func (l *logger) Printf(format string, v ...interface{}) {
	l.l.Info(fmt.Sprintf(format, v...))
	l.l.Sync()
}

func (l *logger) Print(v ...interface{}) {
	l.l.Info(fmt.Sprintf("%v", v))
	l.l.Sync()
}

func GetLogger() BookstoreLogger {
	return &log
}
