package sort

import (
	"testing"

	"github.com/colinfaivre/go-dsa/utils"
)

func TestSelectionSort(t *testing.T) {
	arr, _ := utils.ReadIntegersFromFile("../data/100_000_numbers")
	received := SelectionSort(arr)

	if received[0] != 1 {
		t.Errorf("SelectionSort()[0] expected 0 got %v", received[0])
	}

	if received[99999] != 100000 {
		t.Errorf("SelectionSort()[99999] expected 100000 got %v", received[99999])
	}
}
