package problems

import "sort"

type Job [2]int

func SortByDiff(jobs [][2]int) [][2]int {
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i][0]-jobs[i][1] < jobs[j][0]-jobs[j][1] {
			return false
		} else if jobs[i][0]-jobs[i][1] > jobs[j][0]-jobs[j][1] {
			return true
		} else {
			return jobs[i][0] > jobs[j][0]
		}
	})

	return jobs
}

func SortByRatio(jobs [][2]int) [][2]int {
	sort.Slice(jobs, func(i, j int) bool {
		return float64(jobs[i][0])/float64(jobs[i][1]) > float64(jobs[j][0])/float64(jobs[j][1])
	})

	return jobs
}

func CompletionTime(jobs [][2]int) int {
	var result int
	var time int

	for _, job := range jobs {
		time += job[1]
		result += time * job[0]
	}

	return result
}
