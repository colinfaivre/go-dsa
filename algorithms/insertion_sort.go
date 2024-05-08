package algorithms

// O(n^2)
// The insertion sort algorithm builds the final sorted array one value at a time.
// It assumes that the first element is already sorted.
// Then, a comparison with the second value is performed
// Then the comparison will take place with the third value to insert it in first, second or third position, and so on.
func InsertionSort(num_list []int) {
	n := len(num_list)

	for i := 1; i < n; i++ {
		j := i
		temp := num_list[i]

		for j > 0 && num_list[j-1] > temp {
			num_list[j] = num_list[j-1]
			j--
		}

		num_list[j] = temp
	}

}
