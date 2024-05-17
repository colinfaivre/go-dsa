package algorithms_test

import (
	"github.com/colinfaivre/go-dsa/algorithms"
	"github.com/colinfaivre/go-dsa/parsing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HeapSort", func() {
	Context("Sort a 100_000 unique integers array with values from 1 to 100_000", func() {
		arr, _ := parsing.ReadIntegersFromFile("../test/data/100_000_numbers")
		originalArrLength := len(arr)
		// fmt.Println(time.Now().Unix())
		sortedArr := algorithms.HeapSort(arr)
		// fmt.Println(time.Now().Unix())

		It("has 1 as first value", func() {
			Expect(sortedArr[0]).To(Equal(1))
		})

		It("has 100_000 as last value", func() {
			Expect(sortedArr[len(sortedArr)-1]).To(Equal(100_000))
		})

		It("has the same length as the original array", func() {
			Expect(len(sortedArr)).To(Equal(originalArrLength))
		})

	})
})
