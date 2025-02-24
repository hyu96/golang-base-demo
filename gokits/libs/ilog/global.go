package ilog

import (
	"context"

	iconst "github.com/huydq/gokits/constants"
)

var global *Logger

func SetOptions(opts ...Option) {
	global = global.WithOptions(opts...)
}

func Debug(args ...interface{}) {
	global.debug(global.base, args...)
}

func Debugf(format string, args ...interface{}) {
	global.debugf(global.base, format, args...)
}

func Debugln(args ...interface{}) {
	global.debug(global.base, sprintln(args...))
}

func Debugw(msg string, keysAndValues ...interface{}) {
	global.debugw(global.base, msg, keysAndValues...)
}

func Info(args ...interface{}) {
	global.info(global.base, args...)
}

func Infof(format string, args ...interface{}) {
	global.infof(global.base, format, args...)
}

func Infoln(args ...interface{}) {
	global.info(global.base, sprintln(args...))
}

func Infow(msg string, keysAndValues ...interface{}) {
	global.infow(global.base, msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	global.warn(global.base, args...)
}

func Warnf(template string, args ...interface{}) {
	global.warnf(global.base, template, args...)
}

func Warnln(args ...interface{}) {
	global.warn(global.base, sprintln(args...))
}

func Warnw(msg string, keysAndValues ...interface{}) {
	global.warnw(global.base, msg, keysAndValues...)
}

func Error(args ...interface{}) {
	global.error(global.base, args...)
}

func Errorf(template string, args ...interface{}) {
	global.errorf(global.base, template, args...)
}

func Errorln(args ...interface{}) {
	global.error(global.base, sprintln(args...))
}

func Errorw(msg string, keysAndValues ...interface{}) {
	global.errorw(global.base, msg, keysAndValues...)
}

func DPanic(args ...interface{}) {
	global.dpanic(global.base, args...)
}

func DPanicf(template string, args ...interface{}) {
	global.dpanicf(global.base, template, args...)
}

func DPanicln(args ...interface{}) {
	global.dpanic(global.base, sprintln(args...))
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	global.dpanicw(global.base, msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	global.panic(global.base, args...)
}

func Panicf(template string, args ...interface{}) {
	global.panicf(global.base, template, args...)
}

func Panicln(args ...interface{}) {
	global.panic(global.base, sprintln(args...))
}

func Panicw(msg string, keysAndValues ...interface{}) {
	global.panicw(global.base, msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	global.fatal(global.base, args...)
}

func Fatalf(template string, args ...interface{}) {
	global.fatalf(global.base, template, args...)
}

func Fatalln(args ...interface{}) {
	global.fatal(global.base, sprintln(args...))
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	global.fatalw(global.base, msg, keysAndValues...)
}

// ----------------------------------------------------------------
// with context
func DebugCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Debug(args...)
}

func DebugfCtx(ctx context.Context, format string, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Debugf(format, args...)
}

func DebuglnCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Debugln(sprintln(args...))
}

func DebugwCtx(ctx context.Context, msg string, keysAndValues ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Debugw(msg, keysAndValues...)
}

func InfoCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Info(args...)
}

func InfofCtx(ctx context.Context, format string, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Infof(format, args...)
}

func InfolnCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Infoln(sprintln(args...))
}

func InfowCtx(ctx context.Context, msg string, keysAndValues ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Infow(msg, keysAndValues...)
}

func WarnCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Warn(args...)
}

func WarnfCtx(ctx context.Context, template string, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Warnf(template, args...)
}

func WarnlnCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Warnln(sprintln(args...))
}

func WarnwCtx(ctx context.Context, msg string, keysAndValues ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Warnw(msg, keysAndValues...)
}

func ErrorCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Error(args...)
}

func ErrorfCtx(ctx context.Context, template string, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Errorf(template, args...)
}

func ErrorlnCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Errorln(global.base, sprintln(args...))
}

func ErrorwCtx(ctx context.Context, msg string, keysAndValues ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Errorw(msg, keysAndValues...)
}

func DPanicCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		DPanic(args...)
}

func DPanicfCtx(ctx context.Context, template string, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		DPanicf(template, args...)
}

func DPaniclnCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		DPanic(global.base, sprintln(args...))
}

func DPanicwCtx(ctx context.Context, msg string, keysAndValues ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		DPanicw(msg, keysAndValues...)
}

func PanicCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Panic(args...)
}

func PanicfCtx(ctx context.Context, template string, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Panicf(template, args...)
}

func PaniclnCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Panicln(sprintln(args...))
}

func PanicwCtx(ctx context.Context, msg string, keysAndValues ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Panicw(msg, keysAndValues...)
}

func FatalCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Fatal(args...)
}

func FatalfCtx(ctx context.Context, template string, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Fatalf(template, args...)
}

func FatallnCtx(ctx context.Context, args ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Fatalln(sprintln(args...))
}

func FatalwCtx(ctx context.Context, msg string, keysAndValues ...interface{}) {
	global.base.
		With(string(iconst.KContextKeyRequestID), ctx.Value(iconst.KContextKeyRequestID).(string)).
		Fatalw(msg, keysAndValues...)
}

func With(keysAndValues ...interface{}) *Logger {
	childLogger := *global
	childLogger.base = childLogger.base.With(keysAndValues...)
	return &childLogger
}

func Rotate() error {
	return global.Rotate()
}
