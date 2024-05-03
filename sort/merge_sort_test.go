package sort

import (
	"testing"

	"github.com/colinfaivre/go-dsa/utils"
)

// @TODO white huge slice to file for comparison
func TestMergeSort(t *testing.T) {
	arr, _ := utils.ReadIntegersFromFile("../data/100_000_numbers")
	received := MergeSort(arr)

	if received[0] != 1 {
		t.Errorf("MergeSort()[0] expected 0 got %v", received[0])
	}

	if received[99999] != 100000 {
		t.Errorf("MergeSort()[99999] expected 100000 got %v", received[99999])
	}

	// for i := range arr {
	// 	if expected[i] != received[i] {
	// 		t.Errorf("MergeSort() expected %v got %v", expected, received)
	// 	}
	// }

	// t.Logf("log %v", received)
}
