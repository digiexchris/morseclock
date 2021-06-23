package timing_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTiming(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Timing Suite")
}
