package problems_test

import (
	"strconv"

	"github.com/colinfaivre/go-dsa/parsing"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("clustering", func() {
	Context("SwicthBit()", func() {
		It("Switches a bit from 1 to 0 at a given position", func() {
			bin_num, _ := strconv.ParseUint("1111", 2, 32)
			expected, _ := strconv.ParseUint("1110", 2, 32)

			Expect(problems.SwitchBit(bin_num, 0)).To(Equal(expected))
		})

		It("Switches a bit from 0 to 1 at a given position", func() {
			bin_num, _ := strconv.ParseUint("0000", 2, 32)
			expected, _ := strconv.ParseUint("0001", 2, 32)

			Expect(problems.SwitchBit(bin_num, 0)).To(Equal(expected))
		})
	})

	Context("Huge array", func() {
		It("returns a binary numbers slice", func() {
			arr, _ := parsing.ReadBinIntegersFromFile("../test/data/clustering_bg")

			Expect(arr[0]).To(Equal(uint64(14734287)))
		})
	})
})
