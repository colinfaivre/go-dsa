package problems

func kkmax(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Knapsack(items [][2]int, w, n int) int {
	A := make([][]int, n+1)
	for i := range A {
		A[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for x := 1; x <= w; x++ {
			if items[i-1][1] > x {
				A[i][x] = A[i-1][x]
			} else {
				A[i][x] = kkmax(A[i-1][x], A[i-1][x-items[i-1][1]]+items[i-1][0])
			}
		}
	}

	return A[n][w]
}

// @MEDIUM helpful article https://medium.com/swlh/0-1-knapsack-problem-memoized-day-42-python-f0562f38293e
func RecursiveKnapsack(items [][2]int, w, n int) int {
	if n == 0 || w == 0 {
		return 0
	}

	if items[n-1][1] > w {
		return RecursiveKnapsack(items, w, n-1)
	} else {
		return kkmax(items[n-1][0]+RecursiveKnapsack(items, w-items[n-1][1], n-1), RecursiveKnapsack(items, w, n-1))
	}
}
