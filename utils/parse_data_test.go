package utils

import (
	"testing"
)

func TestReadIntegersFromFile(t *testing.T) {
	received, err := ReadIntegersFromFile("../data/100_000_numbers")
	if err != nil {
		t.Log(err)
	}

	if received[0] != 54044 {
		t.Errorf("ReadIntegersFromFile(\"../data/100_000_numbers\")[0] expected 54044 but got %v", received[0])
	}

	if received[99999] != 91901 {
		t.Errorf("ReadIntegersFromFile(\"../data/100_000_numbers\")[0] expected 91901 but got %v", received[99999])
	}
}

func TestReadIntegersTuppleFromFile(t *testing.T) {
	received, err := ReadIntegersTuplesFromFile("../data/scc")
	if err != nil {
		t.Log(err)
	}

	t.Logf("tuple %v", received[5105042])

	if received[0][0] != 1 || received[0][1] != 1 {
		t.Errorf("ReadIntegersFromFile(\"../data/scc\")[0] expected [1 1] but got %v", received[0])
	}

	if received[5105042][0] != 875714 || received[5105042][1] != 542453 {
		t.Errorf("ReadIntegersFromFile(\"../data/scc\")[0] expected [875714 542453] but got %v", received[5105042])
	}
}
