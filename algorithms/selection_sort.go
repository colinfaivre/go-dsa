package algorithms

// WIKI https://en.wikipedia.org/wiki/Selection_sort

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		indexMin := i

		for j := i + 1; j < len(arr); j++ {
			if arr[indexMin] > arr[j] {
				indexMin = j
			}
		}

		if indexMin != i {
			temp := arr[i]
			arr[i] = arr[indexMin]
			arr[indexMin] = temp
		}
	}

	return arr
}
