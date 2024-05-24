package problems_test

import (
	"github.com/colinfaivre/go-dsa/parsing"
	"github.com/colinfaivre/go-dsa/problems"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Completion time", func() {

	It("computes the right completion time", func() {
		var jobs = [][2]int{{3, 4}, {2, 1}, {4, 6}, {3, 2}, {4, 3}}
		Expect(problems.CompletionTime(jobs)).To(Equal(169))
	})

	It("sorts by diff: when diffs are equal, greater weight goes first", func() {
		jobs := [][2]int{{2, 1}, {3, 2}, {4, 3}, {3, 2}, {4, 3}, {100, 1}}
		problems.SortByDiff(jobs)
		Expect(jobs).To(Equal([][2]int{{100, 1}, {4, 3}, {4, 3}, {3, 2}, {3, 2}, {2, 1}}))
	})

	It("sorts by ratio", func() {
		jobs := [][2]int{{1, 5}, {1, 4}, {1, 3}, {1, 2}, {1, 1}}
		problems.SortByRatio(jobs)
		Expect(jobs).To(Equal([][2]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}}))
	})

	It("computes the assignement right completion time for diffs", func() {
		jobs, _ := parsing.ReadIntegers2TuplesFromFile("../test/data/jobs")
		problems.SortByDiff(jobs)
		Expect(problems.CompletionTime(jobs)).To(Equal(69119377652))
	})

	It("computes the assignement right completion time for ratios", func() {
		jobs, _ := parsing.ReadIntegers2TuplesFromFile("../test/data/jobs")
		problems.SortByRatio(jobs)
		Expect(problems.CompletionTime(jobs)).To(Equal(67311454237))
	})
})
