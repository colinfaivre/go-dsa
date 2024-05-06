package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	arr := []int{3, 2, 1}
	expected := []int{1, 2, 3}
	InsertionSort(arr)

	for i := range arr {
		if expected[i] != arr[i] {
			t.Errorf("InsertionSort() expected %v got %v", expected, arr)
		}
	}
}
