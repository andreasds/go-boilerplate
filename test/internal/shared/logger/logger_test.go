package logger_test

import (
	"testing"

	"github.com/andreasds/go-boilerplate/internal/shared/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	t.Run("SetLogLevel", SetLogLevelTestCase)
}

func SetLogLevelTestCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "success trace",
			input:    "trace",
			expected: "trace",
		},
		{
			name:     "success debug",
			input:    "debug",
			expected: "debug",
		},
		{
			name:     "success info",
			input:    "info",
			expected: "info",
		},
		{
			name:     "success warn",
			input:    "warn",
			expected: "warn",
		},
		{
			name:     "success error",
			input:    "error",
			expected: "error",
		},
		{
			name:     "success fatal",
			input:    "fatal",
			expected: "fatal",
		},
		{
			name:     "success empty default",
			input:    "",
			expected: "trace",
		},
		{
			name:     "success unknown level",
			input:    "unknown",
			expected: "trace",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logger.InitLogger()
			logger.SetLogLevel(test.input)
			got := zerolog.GlobalLevel()

			assert.EqualValues(t, got.String(), test.expected)
		})
	}
}
