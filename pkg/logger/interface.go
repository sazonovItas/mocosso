package logger

type DefaultLogger interface {
	Log(lvl Level, msg string, keysAndValues ...any)
}

type Logger interface {
	DefaultLogger
	Debug(msg string, keysAndValues ...any)
	Info(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
	Error(msg string, keysAndValues ...any)
}

type LogFunc func(lvl Level, msg string, keysAndValues ...any)

func (lf LogFunc) Log(lvl Level, msg string, keysAndValues ...any) {
	lf(lvl, msg, keysAndValues...)
}

func (lf LogFunc) Debug(msg string, keysAndValues ...any) {
	lf(DebugLevel, msg, keysAndValues...)
}

func (lf LogFunc) Info(msg string, keysAndValues ...any) {
	lf(InfoLevel, msg, keysAndValues...)
}

func (lf LogFunc) Warn(msg string, keysAndValues ...any) {
	lf(WarnLevel, msg, keysAndValues...)
}

func (lf LogFunc) Error(msg string, keysAndValues ...any) {
	lf(ErrorLevel, msg, keysAndValues...)
}
