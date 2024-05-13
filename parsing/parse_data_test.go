package parsing_test

import (
	"github.com/colinfaivre/go-dsa/parsing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse Data", func() {
	Context("ReadIntegersFromFile", func() {
		It("does not return an error when the file exists", func() {
			_, err := parsing.ReadIntegersFromFile("../test/data/100_000_numbers")
			Expect(err).To(BeNil())
		})

		It("returns an error when the file does not exist", func() {
			_, err := parsing.ReadIntegersFromFile("../../wrong/path")
			Expect(err).NotTo(BeNil())
		})

		It("returns an array with the correct first and last values", func() {
			arr, _ := parsing.ReadIntegersFromFile("../test/data/100_000_numbers")
			Expect(arr[0]).To(Equal(54044))
			Expect(arr[99999]).To(Equal(91901))
		})
	})

	Context("ReadIntegersTuplesFromFile", func() {
		It("does not return an error when the file exists", func() {
			_, err := parsing.ReadIntegersTuplesFromFile("../test/data/directed_graph")
			Expect(err).To(BeNil())
		})

		It("returns an error when the file does not exist", func() {
			_, err := parsing.ReadIntegersTuplesFromFile("../wrong/path")
			Expect(err).NotTo(BeNil())
		})

		It("returns an array of tuples with the correct first and last tuples", func() {
			arr, _ := parsing.ReadIntegersTuplesFromFile("../test/data/directed_graph")
			Expect(arr[0]).To(Equal([2]int{1, 1}))
			Expect(arr[5105042]).To(Equal([2]int{875714, 542453}))
		})
	})
})
