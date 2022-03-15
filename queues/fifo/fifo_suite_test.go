package fifo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSimpleQueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fifo Suite")
}
