package problems

import (
	"github.com/colinfaivre/go-dsa/datastructures"
)

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

// Takes a binary number and returns a list of all possible binary numbers
// with hamming distance 0, 1 or 2
func ComplementList(num uint64) []uint64 {
	c_list := []uint64{num} // initiate list with 0 hamming distance num
	msb_pos := MostSignificantBitPosition(num)

	// add 1 hamming distance numbers to the map
	for pos := 0; pos <= msb_pos; pos++ {
		complement_num := SwitchBit(num, pos)
		c_list = append(c_list, complement_num)
	}

	// add 2 hamming distance numbers to the map
	for pos1 := 0; pos1 <= msb_pos; pos1++ {
		complement_num := SwitchBit(num, pos1)

	Inner:
		for pos2 := 0; pos2 <= msb_pos; pos2++ {
			complement_num_copy := SwitchBit(complement_num, pos2)
			for _, v := range c_list {
				if v == complement_num_copy {
					continue Inner
				}
			}
			c_list = append(c_list, complement_num_copy)
		}
	}

	return c_list
}

func GetNumToIdx(nums []uint64) map[uint64][]int {
	numToIdx := map[uint64][]int{}

	for i, num := range nums {
		_, _ok := numToIdx[num]
		if !_ok {
			numToIdx[num] = []int{i + 1}
		} else {
			numToIdx[num] = append(numToIdx[num], i+1)
		}
	}

	return numToIdx
}

func GetAdjList(nums []uint64) [][]uint64 {
	adj_list := [][]uint64{}

	v_map := map[uint64]bool{}
	for _, k := range nums {
		v_map[k] = true
	}

	for _, num := range nums {
		edge_list := []uint64{}
		for _, v := range ComplementList(num) {
			if v_map[v] {
				edge_list = append(edge_list, v)
			}
		}

		adj_list = append(adj_list, edge_list)
	}

	return adj_list
}

func GetClusteringResult(nums []uint64) int {
	adj_list := GetAdjList(nums)
	cluster_counter := len(nums)
	num_to_idx := GetNumToIdx(nums)

	uf := datastructures.NewUnionFind(20_000_000)

	for vi, v := range adj_list {
		for _, w := range v {
			for _, wi := range num_to_idx[w] {
				if uf.Union(vi+1, wi) {
					cluster_counter--
				}
			}
		}
	}

	return cluster_counter
}
