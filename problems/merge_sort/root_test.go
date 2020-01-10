package mergeSort

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

// TextLoggerSetup will test the logging setup to make sure it doesnt report
// any unexpected errors for common setup
func TestLoggerSetup(t *testing.T) {

	RegisterFailHandler(Fail)    // registers the fail handler from ginkgo
	RunSpecs(t, "Logging setup") // hands over control to the ginkgo testing framework
}

// Describe runs the testing
var _ = Describe("merge sort", func() {
	Context("sort array", func() {
		It("will be in increasing order", func() {
			array := []int{23, 12, 1, 24, 7, 9, 19}
			result := Sort(array)
			Expect(result).To(Equal([]int{1, 7, 9, 12, 19, 23, 24}))
		})
	})

	Context("if given an empty array", func() {
		It("nil", func() {
			array := []int{}
			result := Sort(array)
			Expect(result).To(BeNil())
		})
	})

	Context("if given an empty array", func() {
		It("nil", func() {
			array := []int{}
			result := Sort(array)
			Expect(result).To(BeNil())
		})
	})

	Context("if given an empty array", func() {
		It("nil", func() {
			array := []int{}
			result := Sort(array)
			Expect(result).To(BeNil())
		})
	})
})
