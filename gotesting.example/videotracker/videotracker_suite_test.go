package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVideotracker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Videotracker Suite")
}
