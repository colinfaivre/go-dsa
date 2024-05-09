package parsing

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParseData(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Suite")
}

var _ = Describe("Parse Data", func() {
	Context("ReadIntegersFromFile", func() {
		It("does not return an error when the file exists", func() {
			_, err := ReadIntegersFromFile("../test/data/100_000_numbers")
			Expect(err).To(BeNil())
		})

		It("returns an error when the file does not exist", func() {
			_, err := ReadIntegersFromFile("../../wrong/path")
			Expect(err).NotTo(BeNil())
		})

		It("returns an array with the correct first and last values", func() {
			arr, _ := ReadIntegersFromFile("../test/data/100_000_numbers")
			Expect(arr[0]).To(Equal(54044))
			Expect(arr[99999]).To(Equal(91901))
		})
	})

	Context("ReadIntegersTuplesFromFile", func() {
		It("does not return an error when the file exists", func() {
			_, err := ReadIntegersTuplesFromFile("../test/data/scc")
			Expect(err).To(BeNil())
		})

		It("returns an error when the file does not exist", func() {
			_, err := ReadIntegersTuplesFromFile("../wrong/path")
			Expect(err).NotTo(BeNil())
		})

		It("returns an array of tuples with the correct first and last tuples", func() {
			arr, _ := ReadIntegersTuplesFromFile("../test/data/scc")
			Expect(arr[0]).To(Equal([2]int{1, 1}))
			Expect(arr[5105042]).To(Equal([2]int{875714, 542453}))
		})
	})
})
