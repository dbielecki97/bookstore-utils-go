package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log = logger{}
)

type logger struct {
	l *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
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

func GetLogger() *logger {
	return &log
}
