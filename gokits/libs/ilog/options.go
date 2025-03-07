package ilog

import (
	"io"

	"go.uber.org/zap/zapcore"
)

// =====================================================

type RotationConfig struct {
	MaxSize    int  `json:"maxSize"`    // megabytes
	MaxAge     int  `json:"maxAge"`     // days
	MaxBackups int  `json:"maxBackups"` // count
	LocalTime  bool `json:"localTime"`
	Compress   bool `json:"compress"`
}

func (c RotationConfig) apply(o *options) {
	if c.MaxAge > 0 {
		o.MaxAge = c.MaxAge
	}

	if c.MaxBackups > 0 {
		o.MaxBackups = c.MaxBackups
	}

	if c.MaxSize > 0 {
		o.MaxSize = c.MaxSize
	}

	o.Compress = c.Compress
	o.LocalTime = c.LocalTime
}

// =====================================================

type options struct {
	RotationConfig

	LogFileLevel    Level
	LogConsoleLevel Level
	Format          Format
	Encoder         zapcore.Encoder

	Development bool

	Output   io.Writer
	LogDirs  []string
	LogFiles []string

	AddCaller  bool
	CallerSkip int
}

func (o options) Clone() options {
	c := options{
		RotationConfig:  o.RotationConfig,
		LogFileLevel:    o.LogFileLevel,
		LogConsoleLevel: o.LogConsoleLevel,
		Format:          o.Format,

		Development: o.Development,

		Output: o.Output,

		AddCaller:  o.AddCaller,
		CallerSkip: o.CallerSkip,
	}

	if o.Encoder != nil {
		c.Encoder = o.Encoder.Clone()
	}

	if len(o.LogDirs) > 0 {
		c.LogDirs = make([]string, len(o.LogDirs))

		copy(c.LogDirs, o.LogDirs)
	}

	if len(o.LogFiles) > 0 {
		c.LogFiles = make([]string, len(o.LogFiles))

		copy(c.LogFiles, o.LogFiles)
	}

	return c
}

func (o options) ZapFileLevelEnabled(lvl zapcore.Level) bool {
	return o.Development || o.LogFileLevel.Enabled(fromZapLevel(lvl))
}

func (o options) ZapConsoleLevelEnabled(lvl zapcore.Level) bool {
	return o.Development || o.LogConsoleLevel.Enabled(fromZapLevel(lvl))
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(log *options) {
	f(log)
}

func WithFileLevel(lvl Level) Option {
	return optionFunc(func(l *options) {
		l.LogFileLevel = lvl
	})
}

func WithConsoleLevel(lvl Level) Option {
	return optionFunc(func(l *options) {
		l.LogConsoleLevel = lvl
	})
}

func WithFormat(format Format) Option {
	return optionFunc(func(l *options) {
		l.Format = format
	})
}

func Development() Option {
	return WithDevelopment(true)
}

func WithDevelopment(development bool) Option {
	return optionFunc(func(l *options) {
		l.Development = development
	})
}

func WithEncoder(encoder zapcore.Encoder) Option {
	return optionFunc(func(l *options) {
		l.Encoder = encoder
	})
}

func WithOutput(output io.Writer) Option {
	return optionFunc(func(l *options) {
		l.Output = output
	})
}

func WithLogDirs(dirs ...string) Option {
	return optionFunc(func(l *options) {
		dst := make([]string, len(dirs))
		copy(dst, dirs)

		l.LogDirs = dst
	})
}

func WithLogFiles(files ...string) Option {
	return optionFunc(func(l *options) {
		dst := make([]string, len(files))
		copy(dst, files)

		l.LogFiles = dst
	})
}

func WithCaller(caller bool) Option {
	return optionFunc(func(l *options) {
		l.AddCaller = caller
	})
}

func AddCaller() Option {
	return WithCaller(true)
}

func AddCallerSkip(skip int) Option {
	return optionFunc(func(l *options) {
		l.CallerSkip += skip
	})
}
