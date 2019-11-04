package log

import (
	"context"
	"fmt"

	"github.com/tron-us/go-common/constant"
	"github.com/tron-us/go-common/env"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger

	// Expose methods under the package for shorter reference
	Debug func(string, ...zap.Field)
	Info  func(string, ...zap.Field)
	Warn  func(string, ...zap.Field)
	Error func(string, ...zap.Field)
	Panic func(string, ...zap.Field)
	With  func(...zap.Field) *zap.Logger
)

func init() {
	var cfg zap.Config
	if env.IsDev() {
		cfg = zap.NewDevelopmentConfig()
	} else {
		cfg = zap.NewProductionConfig()
	}

	if env.LogFile != "" {
		cfg.OutputPaths = append(cfg.OutputPaths, env.LogFile)
	}

	var err error
	logger, err = cfg.Build()
	if err != nil {
		// No logger available, use plain-old panic
		panic(fmt.Sprintf("%v [%v]", constant.LogInitError, err))
	}

	Debug = logger.Debug
	Info = logger.Info
	Warn = logger.Warn
	Error = logger.Error
	Panic = logger.Panic
	With = logger.With
}

// LogWarnErr logs an error at warn level and returns the same message
func HandlerWarnErr(ctx context.Context, msg string, fields ...zap.Field) error {
	err := fmt.Errorf(msg)
	HandlerWarn(ctx, msg, fields...)
	return err
}

// SetLoggerContextFields sets a custom zap.Field into the current context to be
// recorded by the logger when used, and returns new context
func SetLoggerContextFields(ctx context.Context, fields ...zap.Field) context.Context {
	if ctx == nil {
		ctx = context.Background() // create empty context by default
	}
	var logFields []zap.Field
	if logFieldKey := ctx.Value(constant.ContextLogFieldKey); logFieldKey != nil {
		logFields = logFieldKey.([]zap.Field)
	}
	// Append to end of current log fields
	return context.WithValue(ctx, constant.ContextLogFieldKey, append(logFields, fields...))
}

func setLoggerContext(ctx context.Context) *zap.Logger {
	loggerTmp := logger
	if ctx != nil {
		if handlerKey := ctx.Value(constant.HandlerNameContext); handlerKey != nil {
			loggerTmp = loggerTmp.Named(handlerKey.(string))
		}
		if urlKey := ctx.Value(constant.HTTPURLContext); urlKey != nil {
			loggerTmp = loggerTmp.Named(urlKey.(string))
		}
		if logFieldKey := ctx.Value(constant.ContextLogFieldKey); logFieldKey != nil {
			loggerTmp = loggerTmp.With(logFieldKey.([]zap.Field)...)
		}
	}
	return loggerTmp
}

// wrap error with handler name
func HandlerError(ctx context.Context, msg string, fields ...zap.Field) {
	setLoggerContext(ctx).Error(msg, fields...)
}

// wrap warn with handler name
func HandlerWarn(ctx context.Context, msg string, fields ...zap.Field) {
	setLoggerContext(ctx).Warn(msg, fields...)
}

// wrap debug with handler name
func HandlerDebug(ctx context.Context, msg string, fields ...zap.Field) {
	setLoggerContext(ctx).Debug(msg, fields...)
}

// wrap info with handler name
func HandlerInfo(ctx context.Context, msg string, fields ...zap.Field) {
	setLoggerContext(ctx).Info(msg, fields...)
}

// wrap error with handler name
func HandlerPanic(ctx context.Context, msg string, fields ...zap.Field) {
	setLoggerContext(ctx).Panic(msg, fields...)
}
