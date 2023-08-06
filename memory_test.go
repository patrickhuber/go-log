package log_test

import (
	"bytes"
	"testing"

	"github.com/patrickhuber/go-log"
	"github.com/stretchr/testify/require"
)

func TestMemory(t *testing.T) {
	t.Run("log when debug set", func(t *testing.T) {
		var out bytes.Buffer
		logger := log.MemoryWith(&out, log.SetLevel(log.DebugLevel))
		logger.Debug("test")
		require.NotEqual(t, 0, out.Len())
	})
	t.Run("does not log when error set", func(t *testing.T) {
		var out bytes.Buffer
		logger := log.MemoryWith(&out, log.SetLevel(log.ErrorLevel))
		logger.Debug("test")
		require.Equal(t, 0, out.Len())
	})
	t.Run("level", func(t *testing.T) {
		logger := log.Memory(log.SetLevel(log.FatalLevel))
		require.Equal(t, log.FatalLevel, logger.Level())
	})
}
