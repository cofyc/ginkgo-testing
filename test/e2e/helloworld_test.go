package e2e

import (
	"fmt"

	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("Hello World Test", func() {
	ginkgo.BeforeEach(func() {
		ginkgo.By("Before hello world tests")
	})

	ginkgo.AfterEach(func() {
		ginkgo.By("After hello world tests")
	})

	ginkgo.It("a simple test", func() {
		ginkgo.By("run a simple test")
	})

	// Generate test specs programmally.
	ginkgo.Context("matrix tests", func() {
		ginkgo.BeforeEach(func() {
			ginkgo.By("Before matrix tests")
		})

		ginkgo.AfterEach(func() {
			ginkgo.By("After matrix tests")
		})

		versions := []string{"v1", "v2", "v3"}
		for _, v := range versions {
			localV := v // new local variable, otherwise it functions may reference the same variable in 'for range' loop
			ginkgo.It(fmt.Sprintf("test %v", localV), func() {
				ginkgo.By(fmt.Sprintf("testing with version %v", localV))
			})
		}
	})
})
