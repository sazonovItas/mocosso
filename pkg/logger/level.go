package logger

import (
	"strings"
)

// A Level is a logging priority level.
type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

const (
	debugLvl string = "debug"
	infoLvl  string = "info"
	warnLvl  string = "warn"
	errorLvl string = "error"
	panicLvl string = "panic"
	fatalLvl string = "fatal"
)

func LevelFromString(s string) Level {
	s = strings.ToLower(s)

	switch s {
	case debugLvl:
		return DebugLevel
	case infoLvl:
		return InfoLevel
	case warnLvl:
		return WarnLevel
	case errorLvl:
		return ErrorLevel
	case panicLvl:
		return PanicLevel
	case fatalLvl:
		return FatalLevel
	}

	return DebugLevel
}

// String method returns ASCII representation of the log level.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return debugLvl
	case InfoLevel:
		return infoLvl
	case WarnLevel:
		return warnLvl
	case ErrorLevel:
		return errorLvl
	case PanicLevel:
		return panicLvl
	case FatalLevel:
		return fatalLvl
	}

	return debugLvl
}

// CapitalString method returns capital ASCII representation of the log level.
func (l Level) CapitalString() string {
	return strings.ToUpper(l.String())
}
