package log_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/patrickhuber/go-log"
)

var _ = Describe("Memory", func() {
	var (
		out bytes.Buffer
	)
	Describe("Debug", func() {
		It("can log when debug set", func() {
			logger := log.MemoryWith(&out, log.SetLevel(log.DebugLevel))
			logger.Debug("test")
			Expect(out.Len()).ToNot(BeZero())
		})
		It("does not log when error set", func() {
			logger := log.MemoryWith(&out, log.SetLevel(log.ErrorLevel))
			logger.Debug("test")
			Expect(out.Len()).To(BeZero())
		})
	})
	Describe("Level", func() {
		It("can set", func() {
			logger := log.Memory(log.SetLevel(log.FatalLevel))
			Expect(logger.Level()).To(Equal(log.FatalLevel))
		})
	})
	AfterEach(func() {
		out.Reset()
	})
})
