package problems_test

import (
	"github.com/colinfaivre/go-dsa/parsing"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MaxWeightIndependantSet()", func() {
	Context("With small data set", func() {
		weight_list := []int{10, 25, 1, 1, 50, 60}
		max_weight_idependant_set, max_weight := problems.MaxWeightIndependantSet(weight_list)

		It("returns the right max_weight", func() {
			Expect(max_weight).To(Equal(86))
		})

		It("returns the right result", func() {
			result := [6]bool{}
			result[0] = max_weight_idependant_set[1]
			result[1] = max_weight_idependant_set[2]
			result[2] = max_weight_idependant_set[3]
			result[3] = max_weight_idependant_set[4]
			result[4] = max_weight_idependant_set[5]
			result[5] = max_weight_idependant_set[6]

			Expect(result).To(Equal([6]bool{false, true, false, true, false, true}))
		})
	})

	Context("With huge data set", func() {
		weight_list, _ := parsing.ReadIntegersFromFile("../test/data/max_wt_indpdt_set")
		max_weight_idependant_set, max_weight := problems.MaxWeightIndependantSet(weight_list)

		It("returns the right max_weight", func() {
			Expect(max_weight).To(Equal(2955353732))
		})

		It("returns the right result", func() {
			result := [8]bool{}
			result[0] = max_weight_idependant_set[1]
			result[1] = max_weight_idependant_set[2]
			result[2] = max_weight_idependant_set[3]
			result[3] = max_weight_idependant_set[4]
			result[4] = max_weight_idependant_set[17]
			result[5] = max_weight_idependant_set[117]
			result[6] = max_weight_idependant_set[517]
			result[7] = max_weight_idependant_set[997]

			Expect(result).To(Equal([8]bool{true, false, true, false, false, true, true, false}))
		})
	})
})
