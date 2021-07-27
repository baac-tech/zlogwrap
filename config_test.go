package zlogwrap

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

type testCaseDefault struct {
	name     string
	input    interface{}
	expected interface{}
}

func TestConfigDefault(t *testing.T) {
	tests := []testCaseDefault{
		{
			name:     "Empty Input",
			input:    nil,
			expected: &ConfigDefault,
		},
		{
			name:     "With an Input",
			input:    Config{Hidden: true, ServiceName: "test"},
			expected: &Config{Hidden: true, ServiceName: "test", Logger: log.Logger},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.input != nil {
				actual := configDefault(test.input.(Config))
				require.Equal(t, test.expected.(*Config), actual)
				return
			}
			actual := configDefault()
			require.Equal(t, test.expected.(*Config), actual)
		})
	}
}

func TestZLogWrapDefault(t *testing.T) {
	tests := []testCaseDefault{
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
		{
			name:     "With text",
			input:    "testing logger",
			expected: "testing logger",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := configDefault(Config{Logger: log.Output(&buf)})

			logger.Debug(test.input.(string))
			require.Contains(t, buf.String(), test.expected.(string))

			logger.Info(test.input.(string))
			require.Contains(t, buf.String(), test.expected.(string))

			logger.Warn(test.input.(string))
			require.Contains(t, buf.String(), test.expected.(string))

			logger.Error(test.input.(string))
			require.Contains(t, buf.String(), test.expected.(string))

			// logger.Fatal(test.input.(string))                         // TODO: I'll do it later XD
			// require.Contains(t, buf.String(), test.expected.(string)) // TODO: I'll do it later XD

			panicMaker := func() {
				logger.Panic(test.input.(string))
				require.Contains(t, buf.String(), test.expected.(string))
			}
			require.Panics(t, panicMaker)
		})
	}
}

func TestZLogWrapHiddenTrue(t *testing.T) {
	tests := []testCaseDefault{
		{
			name:     "Empty",
			input:    "",
			expected: "",
		},
		{
			name:     "With text",
			input:    "testing logger",
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// actualDebugOut := captureOutput(func() {
			// 	logger := Config{Hidden: true}
			// 	logger.Debug(test.input.(string))
			// })
			// require.Contains(t, actualDebugOut, test.expected.(string))

			// actualInfoOut := captureOutput(func() {
			// 	logger := Config{Hidden: true}
			// 	logger.Info(test.input.(string))
			// })
			// require.Contains(t, actualInfoOut, test.expected.(string))

			// actualErrorOut := captureOutput(func() {
			// 	logger := Config{Hidden: true}
			// 	logger.Error(test.input.(string))
			// })
			// require.Contains(t, actualErrorOut, test.expected.(string))
		})
	}
}
