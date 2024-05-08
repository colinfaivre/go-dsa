package algorithms

import (
	"testing"

	"github.com/colinfaivre/go-dsa/parsing"
)

func TestQuickSort(t *testing.T) {
	arr, _ := parsing.ReadIntegersFromFile("../testdata/100_000_numbers")
	received := QuickSort(arr)

	if received[0] != 1 {
		t.Errorf("QuickSort()[0] expected 0 got %v", received[0])
	}

	if received[99999] != 100000 {
		t.Errorf("QuickSort()[99999] expected 100000 got %v", received[99999])
	}
}
