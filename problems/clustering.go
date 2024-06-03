package problems

import "fmt"

// @COURSERA-DISCUSSION https://www.coursera.org/learn/algorithms-greedy/discussions/forums/N2idJHblEeag2QpBph2LIw/threads/iqd84B0MEe-EcRJB-20IWw
// @COURSERA-DISCUSSION https://www.coursera.org/learn/algorithms-greedy/discussions/forums/N2idJHblEeag2QpBph2LIw/threads/08WOB8G6Eey8MQpu6643JQ

func Clustering(nums []uint64) {
	fmt.Printf("binary number: %024b", nums[0])
}

// Switches a bit from 0 to 1 or 1 to 0 at a given position
func SwitchBit(bin_num uint64, pos int) uint64 {
	bin_num ^= 1 << pos

	return bin_num
}

// Returns the number most significant bit position
func MostSignificantBitPosition(n uint64) int {
	var bit_pos int = 0
	for n != 0 {
		bit_pos++
		n = n >> 1
	}
	return bit_pos - 1
}

// Takes a binary number and returns a map of all possible binary numbers
// with hamming distance 0, 1 or 2
func ComplementMap(num uint64) map[uint64]bool {
	c_map := map[uint64]bool{num: true} // initiate map with 0 hamming distance num
	msb_pos := MostSignificantBitPosition(num)

	// add 1 hamming distance numbers to the map
	for pos := 0; pos <= msb_pos; pos++ {
		complement_num := SwitchBit(num, pos)
		c_map[complement_num] = true
	}

	// add 2 hamming distance numbers to the map
	for pos1 := 0; pos1 <= msb_pos; pos1++ {
		complement_num := SwitchBit(num, pos1)

		for pos2 := 0; pos2 <= msb_pos; pos2++ {
			complement_num_copy := SwitchBit(complement_num, pos2)
			c_map[complement_num_copy] = true
		}
	}

	return c_map
}

// Create a master dictionnary such that
// it takes as key the bit-strings of each vertex and the value is a LIST of vertices with that particular bit-string
// (this is essential to ensure that if there are duplicate bit-strings, they are recognized as separate vertices.
// For example, one of the elements of the dictionary could look like {"110101011...1":[4,5,6,...,7] , ...}
func ComputeMasterDictionnary(nums []uint64) map[uint64][]int {
	master_dict := map[uint64][]int{}

	for _, num := range nums {
		_, ok := master_dict[num]
		if !ok {
			master_dict[num] = []int{}
		}
	}

	for i, num := range nums {
		num_c_map := ComplementMap(num)

		for v := range num_c_map {
			_, _ok := master_dict[v]
			if _ok {
				master_dict[v] = append(master_dict[v], i+1)
			}
		}
	}

	return master_dict
}
