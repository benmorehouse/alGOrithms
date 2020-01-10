package hash_count

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
var _ = Describe("hash count stuff", func() {
	Context("will count the amount of strings in an array", func() {
		It("assertion", func() {
			array := []string{
				"neehoy",
				"menoy",
				"menoy",
				"menoy",
			}
			Expect(Hash(array)).Should(Equal(map[string]int{
				"neehoy": 1,
				"menoy":  3,
			}))
		})
	})

})
