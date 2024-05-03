package utils

import (
	"bufio"
	"os"
	"strconv"
)

// ReadIntegersFromFile reads integers from a file and returns them as a slice.
func ReadIntegersFromFile(filename string) ([]int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var numbers []int

	// Read each line and parse it as an integer
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
