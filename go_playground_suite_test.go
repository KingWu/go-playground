package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoPlayground(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoPlayground Suite")
}
