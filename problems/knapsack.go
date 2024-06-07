package problems

import "strconv"

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

var memo = map[string]int{}

func RecursiveKnapsack(items [][2]int, w, n int) int {
	if n == 0 || w == 0 {
		return 0
	}

	v, ok := memo[strconv.Itoa(w)+"-"+strconv.Itoa(n)]
	if ok {
		return v
	}

	if items[n-1][1] > w {
		memo[strconv.Itoa(w)+"-"+strconv.Itoa(n)] = RecursiveKnapsack(items, w, n-1)
		return memo[strconv.Itoa(w)+"-"+strconv.Itoa(n)]
	} else {
		memo[strconv.Itoa(w)+"-"+strconv.Itoa(n)] = kkmax(items[n-1][0]+RecursiveKnapsack(items, w-items[n-1][1], n-1), RecursiveKnapsack(items, w, n-1))
		return memo[strconv.Itoa(w)+"-"+strconv.Itoa(n)]
	}
}
