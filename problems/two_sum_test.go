package problems_test

import (
	"github.com/colinfaivre/go-dsa/parsing"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("two sum", func() {
	Context("Small array", func() {
		It("returns true if there are two numbers that sum to 11", func() {
			arr := []int{4, 3, 5, 1, 10}
			result := problems.TwoSum(arr, 11)

			Expect(result).To(Equal(true))
		})

		It("returns false if there are not two numbers that sum to 12", func() {
			arr := []int{4, 3, 5, 1, 10}
			result := problems.TwoSum(arr, 12)

			Expect(result).To(Equal(false))
		})
	})

	Context("Huge array", func() {
		It("returns the right number of distinct 'x+y' in the inclusive interval [-10_000, 10_000]", func() {
			huge_arr, _ := parsing.ReadIntegersFromFile("../test/data/two_sum")
			result := problems.GetResult(huge_arr)

			Expect(result).To(Equal(2145))
		})
	})
})
