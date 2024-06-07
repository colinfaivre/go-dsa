package problems_test

import (
	"github.com/colinfaivre/go-dsa/parsing"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Knapsack()", func() {
	Context("small data set", func() {
		It("1 returns the best solution value", func() {
			items, _ := parsing.ReadIntegers2TuplesFromFile("../test/data/knapsack_small")
			knapsack_result := problems.Knapsack(items, 10_000, 100)

			Expect(knapsack_result).To(Equal(2493893))
		})
	})

	Context("big data set", func() {
		It("2 returns the best solution value", func() {
			items, _ := parsing.ReadIntegers2TuplesFromFile("../test/data/knapsack_big")
			knapsack_result := problems.RecursiveKnapsack(items, 2_000_000, 2_000)

			Expect(knapsack_result).To(Equal(2493893))
		})
	})
})
