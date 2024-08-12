package zapper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewZapLogger function returns new zap logger.
func NewZapLogger(
	level zapcore.Level,
	encoding string,
	outputPaths ...string,
) (logger *zap.Logger, err error) {
	encoder := zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		CallerKey:     "caller",
		TimeKey:       "timestamp",
		StacktraceKey: "stacktrace",

		EncodeTime:   zapcore.RFC3339TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          encoding,
		EncoderConfig:     encoder,
		Sampling:          nil,
		OutputPaths:       outputPaths,
		ErrorOutputPaths:  []string{"stderr"},
		// InitialFields: map[string]interface{}{
		// 	"pid": os.Getpid(),
		// },
	}

	return config.Build()
}
