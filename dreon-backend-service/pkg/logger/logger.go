package logger

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	"github.com/hiamthach108/dreon-backend-service/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger  *zap.Logger
	service string
}

func NewLogger(config *config.AppConfig) (ILogger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var level zapcore.Level
	switch config.Logger.Level {
	case string(DebugLv):
		level = zapcore.DebugLevel
	case string(InfoLv):
		level = zapcore.InfoLevel
	case string(WarnLv):
		level = zapcore.WarnLevel
	case string(ErrorLv):
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    encoderConfig,
	}
	zlogger, err := cfg.Build(zap.AddCallerSkip(2))
	if err != nil {
		return nil, err
	}

	return &zapLogger{
		logger:  zlogger,
		service: config.App.Name,
	}, nil
}

func (l *zapLogger) Debug(msg string, fields ...any) {
	l.log(zap.DebugLevel, msg, fields...)
}

func (l *zapLogger) Info(msg string, fields ...any) {
	l.log(zap.InfoLevel, msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...any) {
	l.log(zap.WarnLevel, msg, fields...)
}

func (l *zapLogger) Error(msg string, fields ...any) {
	l.log(zap.ErrorLevel, msg, fields...)
}

func (l *zapLogger) Fatal(msg string, fields ...any) {
	l.log(zap.FatalLevel, msg, fields...)
}

func (l *zapLogger) With(fields ...any) ILogger {
	zapFields := toZapFields(fields...)
	return &zapLogger{
		logger:  l.logger.With(zapFields...),
		service: l.service,
	}
}

func (l *zapLogger) log(level zapcore.Level, msg string, fields ...any) {
	zapFields := make([]zap.Field, 0, len(fields)+1)
	zapFields = append(zapFields, zap.String("service", l.service))
	zapFields = append(zapFields, toZapFields(fields...)...)

	switch level {
	case zap.DebugLevel:
		l.logger.Debug(msg, zapFields...)
	case zap.InfoLevel:
		l.logger.Info(msg, zapFields...)
	case zap.WarnLevel:
		l.logger.Warn(msg, zapFields...)
	case zap.ErrorLevel:
		zapFields = append(zapFields, captureStackTrace())
		l.logger.Error(msg, zapFields...)
	case zap.FatalLevel:
		zapFields = append(zapFields, captureStackTrace())
		l.logger.Fatal(msg, zapFields...)
	}
}

func sanitize(value any) any {
	str, ok := value.(string)
	if !ok {
		jsonBytes, err := json.Marshal(value)
		if err != nil {
			return value
		}
		str = string(jsonBytes)
	}

	strLower := strings.ToLower(str)
	sensitiveFields := []string{
		"password",
		"token",
		"authorization",
		"api_key",
		"secret",
	}
	for _, field := range sensitiveFields {
		if strings.Contains(strLower, field) {
			return "[REDACTED]"
		}
	}

	return value
}

func toZapFields(fields ...any) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields)/2)

	if len(fields)%2 != 0 {
		zapFields = append(zapFields, zap.Any("field_0", fields[0]))
		fields = fields[1:]
	}
	for i := 0; i < len(fields); i += 2 {
		key, ok := fields[i].(string)
		if !ok {
			key = fmt.Sprintf("field_%d", i)
		}
		zapFields = append(zapFields, zap.Any(key, sanitize(fields[i+1])))
	}
	return zapFields
}

func captureStackTrace() zap.Field {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	frames := runtime.CallersFrames(pc)
	var stacktrace string
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		stacktrace += fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
	}
	return zap.String("stacktrace", stacktrace)
}

func (l *zapLogger) GetZapLogger() *zap.Logger {
	return l.logger
}
