package parsing

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

// ReadBinIntegersFromFile reads integers from a file and returns them as a slice.
func ReadBinIntegersFromFile(filename string) ([]uint64, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var numbers []uint64

	// Read each line and parse it as an integer
	for scanner.Scan() {
		line := scanner.Text()
		// remove white spaces
		line = strings.ReplaceAll(line, " ", "")

		number, err := strconv.ParseUint(line, 2, 32)
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

// ReadIntegers2TuplesFromFile reads integers from a file and returns them as a slice of integer tuples.
func ReadIntegers2TuplesFromFile(filename string) ([][2]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numberTuples [][2]int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		first, err_first := strconv.Atoi(split[0])
		second, err_second := strconv.Atoi(split[1])
		if err_first != nil || err_second != nil {
			return nil, err
		}

		numberTuples = append(numberTuples, [2]int{first, second})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numberTuples, nil
}

// ReadIntegers3TuplesFromFile reads integers from a file and returns them as a slice of integer tuples.
func ReadIntegers3TuplesFromFile(filename string) ([][3]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numberTuples [][3]int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		first, err_first := strconv.Atoi(split[0])
		second, err_second := strconv.Atoi(split[1])
		third, err_third := strconv.Atoi(split[2])
		if err_first != nil || err_second != nil || err_third != nil {
			return nil, err
		}

		numberTuples = append(numberTuples, [3]int{first, second, third})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numberTuples, nil
}

func GetWeightedGraphData(filename string) ([][][2]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var result [][][2]int

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Fields(line)
		second := split[1:]
		currentLineArr := [][2]int{}

		for _, e := range second {
			edge := strings.Split(e, ",")
			v, v_err := strconv.Atoi(edge[0])
			w, w_err := strconv.Atoi(edge[1])
			if v_err != nil || w_err != nil {
				return nil, err
			}
			currentLineArr = append(currentLineArr, [2]int{v, w})
		}

		result = append(result, currentLineArr)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
