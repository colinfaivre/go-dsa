package problems

import (
	"fmt"

	"github.com/colinfaivre/go-dsa/datastructures"
)

// @COURSERA-DISCUSSION https://www.coursera.org/learn/algorithms-greedy/discussions/forums/N2idJHblEeag2QpBph2LIw/threads/iqd84B0MEe-EcRJB-20IWw
// @COURSERA-DISCUSSION https://www.coursera.org/learn/algorithms-greedy/discussions/forums/N2idJHblEeag2QpBph2LIw/threads/08WOB8G6Eey8MQpu6643JQ
// @COURSERA-DISCUSSION https://www.coursera.org/learn/algorithms-greedy/discussions/forums/N2idJHblEeag2QpBph2LIw/threads/mSGCyQULEemZUwrw3z8SgA

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
// @T0DO ⚠️ this is essential to ensure that if there are duplicate bit-strings, they are recognized as separate vertices
func ComputeMasterDictionnary(nums []uint64) map[uint64][]uint64 {
	master_dict := map[uint64][]uint64{}

	for _, num := range nums {
		_, ok := master_dict[num]
		if !ok {
			master_dict[num] = []uint64{}
		}
	}

	for _, num := range nums {
		num_c_map := ComplementMap(num)

		for v := range num_c_map {
			_, _ok := master_dict[v]
			if _ok {
				master_dict[v] = append(master_dict[v], num)
			}
		}
	}

	return master_dict
}

func GetClusteringResult(nums []uint64) int {
	not_gathered_counter := 0
	master_dict := ComputeMasterDictionnary(nums)
	cluster_counter := len(master_dict)
	fmt.Println("cluster_counter", cluster_counter) // 198_788 (less than 200_000 so there are duplicates)

	uf := datastructures.NewUnionFind(200_000_000)

	for v, edges := range master_dict {
		for _, w := range edges {
			if uf.Union(int(v), int(w)) {
				cluster_counter--
			}
		}
	}

	for _, edges := range master_dict {
		if len(edges) <= 1 {
			not_gathered_counter++
		}
	}

	return cluster_counter + not_gathered_counter
}
