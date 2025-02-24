package ilog

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

type Level int

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel

	InvalidLevel = _maxLevel + 1
)

func (l Level) Enabled(lvl Level) bool {
	return lvl >= l
}

func ParseLevel(lvl string) Level {
	switch strings.ToLower(lvl) {
	case "fatal":
		return FatalLevel
	case "panic":
		return PanicLevel
	case "dpanic":
		return DPanicLevel
	case "error":
		return ErrorLevel
	case "warn", "warning":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	}

	return InvalidLevel
}

func fromZapLevel(lvl zapcore.Level) Level {
	switch lvl {
	case zapcore.DebugLevel:
		return DebugLevel
	case zapcore.InfoLevel:
		return InfoLevel
	case zapcore.WarnLevel:
		return WarnLevel
	case zapcore.ErrorLevel:
		return ErrorLevel
	case zapcore.DPanicLevel:
		return DPanicLevel
	case zapcore.PanicLevel:
		return PanicLevel
	case zapcore.FatalLevel:
		return FatalLevel
	}

	return InvalidLevel
}

// func toZapLevel(lvl Level) zapcore.Level {
// 	switch lvl {
// 	case DebugLevel:
// 		return zapcore.DebugLevel
// 	case InfoLevel:
// 		return zapcore.InfoLevel
// 	case WarnLevel:
// 		return zapcore.WarnLevel
// 	case ErrorLevel:
// 		return zapcore.ErrorLevel
// 	case DPanicLevel:
// 		return zapcore.DPanicLevel
// 	case PanicLevel:
// 		return zapcore.PanicLevel
// 	case FatalLevel:
// 		return zapcore.FatalLevel
// 	}

// 	return zapcore.InfoLevel
// }
