package core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Core", func() {
	Context("Example context to test core package", func() {
		It("should have an assertion in it", func() {
			Expect(2 + 3).To(Equal(5))
		})
	})

})
