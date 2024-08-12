package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLevelString(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		level    Level
	}{
		{
			name:     "Debug level",
			expected: "debug",
			level:    DebugLevel,
		},
		{
			name:     "Info level",
			expected: "info",
			level:    InfoLevel,
		},
		{
			name:     "Warning level",
			expected: "warning",
			level:    WarnLevel,
		},
		{
			name:     "Error level",
			expected: "error",
			level:    ErrorLevel,
		},
		{
			name:     "Panic level",
			expected: "panic",
			level:    PanicLevel,
		},
		{
			name:     "Fatal level",
			expected: "fatal",
			level:    FatalLevel,
		},
		{
			name:     "Not known level",
			expected: "debug",
			level:    Level(10),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(
				t,
				test.expected,
				test.level.String(),
				"expected %s, got %s",
				test.expected,
				test.level.String(),
			)
		})
	}
}

func TestLevelFromString(t *testing.T) {
	tests := []struct {
		name     string
		level    string
		expected Level
	}{
		{
			name:     "Debug level",
			level:    "debug",
			expected: DebugLevel,
		},
		{
			name:     "Info level",
			level:    "info",
			expected: InfoLevel,
		},
		{
			name:     "Warning level",
			level:    "warning",
			expected: WarnLevel,
		},
		{
			name:     "Error level",
			level:    "error",
			expected: ErrorLevel,
		},
		{
			name:     "Panic level",
			level:    "panic",
			expected: PanicLevel,
		},
		{
			name:     "Fatal level",
			level:    "fatal",
			expected: FatalLevel,
		},
		{
			name:     "Not known level",
			level:    "Level(10)",
			expected: DebugLevel,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(
				t,
				test.expected,
				LevelFromString(test.level),
				"expected %s, got %s",
				test.expected,
				LevelFromString(test.level),
			)
		})
	}
}
