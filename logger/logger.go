package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	logConfig := &zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "time",
			LevelKey:     "level",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	log, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return log
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
	log.Sync()
}

func ErrorWithError(msg string, err error, fields ...zap.Field) {
	fields = append(fields, zap.NamedError("error", err))
	log.Error(msg, fields...)
	log.Sync()
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
	log.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
	log.Sync()
}

func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
	log.Sync()
}

func WarnWithError(msg string, err error, fields ...zap.Field) {
	fields = append(fields, zap.NamedError("warn", err))
	log.Warn(msg, fields...)
	log.Sync()
}

func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
	log.Sync()
}
