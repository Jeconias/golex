package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGolex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golex Suite")
}
