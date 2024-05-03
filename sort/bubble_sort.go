package sort

// O(n^2)
// The bubble sort algorithm compares each every two adjacent values
// and swaps them if the first one is bigger than the second one.
// It has this name because values tend to move up like bubbles rising to the surface.
func BubbleSort(numList []int) {
	for i := len(numList); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numList[j-1] > numList[j] {
				temp := numList[j]
				numList[j] = numList[j-1]
				numList[j-1] = temp
			}
		}
	}
}
