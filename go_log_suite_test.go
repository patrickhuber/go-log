package log_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoLog Suite")
}
