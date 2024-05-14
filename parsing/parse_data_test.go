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

	Context("GetWeightedGraphData", func() {
		It("does not return an error when the file exists", func() {
			_, err := parsing.GetWeightedGraphData("../test/data/directed_weighted_graph")
			Expect(err).To(BeNil())
		})

		It("returns an error when the file does not exist", func() {
			_, err := parsing.GetWeightedGraphData("../wrong/path")
			Expect(err).NotTo(BeNil())
		})

		It("returns correct first and last values", func() {
			arr, _ := parsing.GetWeightedGraphData("../test/data/directed_weighted_graph")
			expected_first := [][2]int{{80, 982}, {163, 8164}, {170, 2620}, {145, 648}, {200, 8021}, {173, 2069}, {92, 647}, {26, 4122}, {140, 546}, {11, 1913}, {160, 6461}, {27, 7905}, {40, 9047}, {150, 2183}, {61, 9146}, {159, 7420}, {198, 1724}, {114, 508}, {104, 6647}, {30, 4612}, {99, 2367}, {138, 7896}, {169, 8700}, {49, 2437}, {125, 2909}, {117, 2597}, {55, 6399}}
			expected_last := [][2]int{{108, 9976}, {103, 6851}, {145, 2753}, {41, 2622}, {187, 6767}, {190, 5999}, {16, 2848}, {194, 2915}, {5, 4009}, {172, 6888}, {39, 4319}, {176, 1709}, {60, 3269}, {138, 678}, {43, 8943}, {98, 2690}, {1, 8021}, {104, 7083}, {154, 229}, {91, 1988}, {67, 475}, {76, 4623}, {195, 8114}, {37, 7541}, {54, 4899}}
			Expect(arr[0]).To(Equal(expected_first))
			Expect(arr[len(arr)-1]).To(Equal(expected_last))
		})
	})
})
