package logger

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestSetLogLevel(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		inputLevel    string
		expectedLevel zerolog.Level
	}{
		{
			name:          "debug",
			inputLevel:    "debug",
			expectedLevel: zerolog.DebugLevel,
		},
		{
			name:          "info",
			inputLevel:    "info",
			expectedLevel: zerolog.InfoLevel,
		},
		{
			name:          "warn",
			inputLevel:    "warn",
			expectedLevel: zerolog.WarnLevel,
		},
		{
			name:          "error",
			inputLevel:    "error",
			expectedLevel: zerolog.ErrorLevel,
		},
		{
			name:          "fallback to InfoLevel",
			inputLevel:    "invalid",
			expectedLevel: zerolog.InfoLevel,
		},
		{
			name:          "fallback to InfoLevel",
			inputLevel:    "",
			expectedLevel: zerolog.InfoLevel,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			SetLogLevel(tc.inputLevel)
			assert.Equal(t, tc.expectedLevel, zerolog.GlobalLevel())
		})
	}
}
