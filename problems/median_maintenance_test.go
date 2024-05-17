package problems_test

import (
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MedianMaintenance()", func() {
	// Context("NaiveMedianMaintenance() on small array", func() {
	// 	arr := []int{4, 1, 7, 3, 2, 5, 9, 6}
	// 	median_list := problems.NaiveMedianMaintenance(arr)
	// 	expected := []int{4, 1, 4, 3, 3, 3, 4, 4}

	// 	It("computes the right median_list", func() {
	// 		Expect(median_list).To(Equal(expected))
	// 	})
	// })

	// Context("NaiveMedianMaintenance() on huge array", func() {
	// 	arr, _ := parsing.ReadIntegersFromFile("../test/data/10_000_numbers")
	// 	median_list := problems.NaiveMedianMaintenance(arr)
	// 	result := problems.GetResult(median_list)

	// 	It("computes the right assignment result", func() {
	// 		Expect(result).To(Equal(1213))
	// 	})
	// })

	Context("HeapMedianMaintenance() on small array", func() {
		arr := []int{4, 1, 7, 3, 2, 5, 9, 6}
		median_list := problems.HeapMedianMaintenance(arr)
		expected := []int{4, 1, 4, 3, 3, 3, 4, 4}

		It("computes the right median_list", func() {
			Expect(median_list).To(Equal(expected))
		})
	})

	// Context("HeapMedianMaintenance() on huge array", func() {
	// 	arr, _ := parsing.ReadIntegersFromFile("../test/data/10_000_numbers")
	// 	median_list := problems.HeapMedianMaintenance(arr)
	// 	result := problems.GetResult(median_list)

	// 	It("computes the right assignment result", func() {
	// 		Expect(result).To(Equal(1213))
	// 	})
	// })
})
