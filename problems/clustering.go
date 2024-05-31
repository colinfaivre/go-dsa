package problems

import "fmt"

// @COURSERA-DISCUSSION https://www.coursera.org/learn/algorithms-greedy/discussions/forums/N2idJHblEeag2QpBph2LIw/threads/iqd84B0MEe-EcRJB-20IWw
// @COURSERA-DISCUSSION https://www.coursera.org/learn/algorithms-greedy/discussions/forums/N2idJHblEeag2QpBph2LIw/threads/08WOB8G6Eey8MQpu6643JQ

func Clustering(nums []uint64) {
	fmt.Printf("binary number: %024b", nums[0])
}

// Switches a bit from 0 to 1 or 1 to 0 at a given position
func SwitchBit(bin_num uint64, pos uint8) uint64 {
	bin_num ^= 1 << pos

	return bin_num
}

// Takes a binary number and returns a set of all possible binary numbers
// with hamming distance 0, 1 or 2
func ComplementSet(num uint64) []uint64 {
	// for input '101', should return:
	// {
	//0
	// 	"101",
	//1
	//	"001",
	// 	"111",
	// 	"100",
	//2
	//	"110",
	// 	"011",
	// 	"000",
	// }
	return []uint64{}
}

// Create a master dictionnary such that
// it takes as key the bit-strings of each vertex and the value is a LIST of vertices with that particular bit-string
// (this is essential to ensure that if there are duplicate bit-strings, they are recognized as separate vertices.
// For example, one of the elements of the dictionary could look like {"110101011...1":[4,5,6,...,7] , ...}
