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

	Context("MostSignificantBitPosition()", func() {
		It("returns the msb of a number", func() {
			bin_num, _ := strconv.ParseUint("1001", 2, 32)

			bin_num2, _ := strconv.ParseUint("1001001", 2, 32)

			Expect(problems.MostSignificantBitPosition(bin_num)).To(Equal(3))
			Expect(problems.MostSignificantBitPosition(bin_num2)).To(Equal(6))
		})
	})

	Context("ComplementMap()", func() {
		It("returns the right complement map", func() {
			bin_num, _ := strconv.ParseUint("101", 2, 32)

			Expect(problems.ComplementMap(bin_num)).To(Equal(map[uint64]bool{0: true, 3: true, 5: true, 4: true, 7: true, 1: true, 6: true}))
		})
	})

	Context("ComputeMasterDictionnary()", func() {
		It("returns the right complement map", func() {
			arr, _ := parsing.ReadBinIntegersFromFile("../test/data/clustering_bg")

			Expect(problems.ComputeMasterDictionnary(arr)[7349808]).To(Equal([]uint64{7415348, 6334000, 7349808, 7349778, 16262704}))
		})
	})

	Context("Huge array", func() {
		It("returns a binary numbers slice", func() {
			arr, _ := parsing.ReadBinIntegersFromFile("../test/data/clustering_bg")

			Expect(arr[0]).To(Equal(uint64(14734287)))
		})
	})

	Context("GetResult()", func() {
		It("returns the right result", func() {
			arr, _ := parsing.ReadBinIntegersFromFile("../test/data/clustering_bg")
			result := problems.GetClusteringResult(arr)

			Expect(result).To(Equal(13306))
		})
	})
})
