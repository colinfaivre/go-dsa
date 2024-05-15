package problems_test

import (
	"github.com/colinfaivre/go-dsa/parsing"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Problems", func() {
	Context("MedianMaintenance() on small array", func() {
		arr := []int{4, 1, 7, 3, 2, 5, 9, 6}
		median_list := problems.MedianMaintenance(arr)
		expected := []int{4, 1, 4, 3, 3, 3, 4, 4}

		It("computes the right median_list", func() {
			Expect(median_list).To(Equal(expected))
		})
	})

	Context("MedianMaintenance() on huge array", func() {
		arr, _ := parsing.ReadIntegersFromFile("../test/data/10_000_numbers")
		median_list := problems.MedianMaintenance(arr)
		result := problems.GetResult(median_list)

		It("computes the right assignment result", func() {
			Expect(result).To(Equal(1213))
		})
	})
})
