package log

import (
	"log"
)

// Logger defines a logging interface
type Logger interface {
	Printf(format string, v ...any)
	Print(v ...any)
	Println(v ...any)
	Fatalf(format string, v ...any)
	Fatal(v ...any)
	Fatalln(v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
	Errorln(v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Warnln(v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Infoln(v ...any)
	Debug(v ...any)
	Debugf(format string, v ...any)
	Debugln(v ...any)
	Trace(v ...any)
	Tracef(format string, v ...any)
	Traceln(v ...any)
	Level() Level
}

// Default returns the default platform logger
func Default(options ...LogOption) Logger {
	l := &logger{
		Logger: log.Default(),
		level:  ErrorLevel,
	}
	for _, opt := range options {
		opt(l)
	}
	return l
}

type logger struct {
	*log.Logger
	level Level
}

func shouldLog(level Level, target Level) bool {
	return level >= target
}

func (l *logger) Error(v ...any) {
	if !shouldLog(l.level, ErrorLevel) {
		return
	}
	l.Print(v...)
}

func (l *logger) Errorf(format string, v ...any) {
	if !shouldLog(l.level, ErrorLevel) {
		return
	}
	l.Printf(format, v...)
}

func (l *logger) Errorln(v ...any) {
	if !shouldLog(l.level, ErrorLevel) {
		return
	}
	l.Println(v...)
}

func (l *logger) Warn(v ...any) {
	if !shouldLog(l.level, WarningLevel) {
		return
	}
	l.Print(v...)
}

func (l *logger) Warnf(format string, v ...any) {
	if !shouldLog(l.level, WarningLevel) {
		return
	}
	l.Printf(format, v...)
}

func (l *logger) Warnln(v ...any) {
	if !shouldLog(l.level, WarningLevel) {
		return
	}
	l.Println(v...)
}

func (l *logger) Info(v ...any) {
	if !shouldLog(l.level, InfoLevel) {
		return
	}
	l.Print(v...)
}

func (l *logger) Infof(format string, v ...any) {
	if !shouldLog(l.level, InfoLevel) {
		return
	}
	l.Printf(format, v...)
}

func (l *logger) Infoln(v ...any) {
	if !shouldLog(l.level, InfoLevel) {
		return
	}
	l.Println(v...)
}

func (l *logger) Debug(v ...any) {
	if !shouldLog(l.level, DebugLevel) {
		return
	}
	l.Print(v...)
}

func (l *logger) Debugf(format string, v ...any) {
	if !shouldLog(l.level, DebugLevel) {
		return
	}
	l.Printf(format, v...)
}

func (l *logger) Debugln(v ...any) {
	if !shouldLog(l.level, DebugLevel) {
		return
	}
	l.Println(v...)
}

func (l *logger) Trace(v ...any) {
	if !shouldLog(l.level, TraceLevel) {
		return
	}
	l.Print(v...)
}

func (l *logger) Tracef(format string, v ...any) {
	if !shouldLog(l.level, TraceLevel) {
		return
	}
	l.Printf(format, v...)
}

func (l *logger) Traceln(v ...any) {
	if !shouldLog(l.level, TraceLevel) {
		return
	}
	l.Println(v...)
}

func (l *logger) Level() Level {
	return l.level
}

type LogOption func(l *logger)

// WithLevel sets the logging level
func WithLevel(level Level) LogOption {
	return func(l *logger) {
		l.level = level
	}
}

// SetLevel sets the logging level
//
// Depricated: use WithLevel instead
func SetLevel(level Level) LogOption {
	return func(l *logger) {
		l.level = level
	}
}
