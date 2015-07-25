package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFactorials(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Factorials Suite")
}
