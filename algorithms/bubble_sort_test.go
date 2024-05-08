package algorithms

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 2, 1}
	expected := []int{1, 2, 3}
	BubbleSort(arr)

	for i := range arr {
		if expected[i] != arr[i] {
			t.Errorf("Bubble_sort() expected %v got %v", expected, arr)
		}
	}
}
