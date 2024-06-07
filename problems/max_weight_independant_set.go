package problems

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MaxWeightIndependantSet(weightList []int) (map[int]bool, int) {
	independantSet := map[int]bool{}

	maxWeightList := []int{}
	maxWeightList = append(maxWeightList, 0)
	maxWeightList = append(maxWeightList, weightList[0])

	for i := 2; i <= len(weightList); i++ {
		maxWeightList = append(maxWeightList, kkmax(maxWeightList[i-1], maxWeightList[i-2]+weightList[i-1]))
	}

	i := len(maxWeightList) - 1

	for i >= 1 {
		var test int
		if i-2 >= 0 {
			test = maxWeightList[i-2]
		}

		if maxWeightList[i-1] >= test+weightList[i-1] {
			i -= 1
		} else {
			independantSet[i] = true
			i -= 2
		}
	}

	return independantSet, maxWeightList[len(maxWeightList)-1]
}
