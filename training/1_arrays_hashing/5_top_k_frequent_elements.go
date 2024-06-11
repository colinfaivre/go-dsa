package training

// 1.5 `M` Top K Frequent Elements

/*** @LEETCODE https://leetcode.com/problems/top-k-frequent-elements/
Given an integer array nums and an integer k,
return the k most frequent elements.
You may return the answer in any order.
---
Example 1:
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
---
Example 2:
Input: nums = [1], k = 1
Output: [1]
---
Constraints:
1 <= nums.length <= 105
-104 <= nums[i] <= 104
k is in the range [1, the number of unique elements in the array].
It is guaranteed that the answer is unique.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=YPTqKIgVk-k
O(nlogn) solution - hashmap/sort
- loop in nums inserting each num into a freqMap - O(n)
- sort the map items in descending order by freqs - O(nlogn)
- pop the first k biggest freq nums - O(k)

O(klogn) solution - hashmap/maxheap
- loop in nums inserting each num into a freqMap - O(n)
- insert all freqMap items into a maxheap - O(n)
- pop the first k biggest frequency nums from the maxheap - O(klogn)

O(n) solution - hashmap/bucketsort
- loop in nums inserting each num into a freqMap - O(n)
- loop in freqMap appending each nums in a subarray at freq index in a freqArr - O(n)
- loop in freqArr in descending order inserting num into the output array
- return output array when its length reaches k
***/

// O(nlogn) solution - hashmap/sort
type Item struct {
    num int
    freq int
}
func topKFrequentOnlogn(nums []int, k int) []int {
    freqMap := map[int]int{}
    freqArr := []Item{}
    result := []int{}

    for _, v := range nums {
        freqMap[v]++
    }

    for key, value := range freqMap {
        freqArr = append(freqArr, Item{num: key, freq: value})
    }

    sort.Slice(freqArr, func(i, j int) bool {
  		return freqArr[i].freq > freqArr[j].freq
  	})

    for i:=1; i <= k; i++ {
        biggestFreq := freqArr[0]
        freqArr = freqArr[1:]
        result = append(result, biggestFreq.num)
    }

    return result
}

// O(n) solution - hashmap/bucketsort
func topKFrequentOn(nums []int, k int) []int {
    freqMap := map[int]int{}
    freqArr := [][]int{}
    for i := 0; i <= len(nums); i++ {
        freqArr = append(freqArr, []int{})
    }

    result := []int{}

    for _, v := range nums {
        freqMap[v]++
    }

    for num, count := range freqMap {
        freqArr[count] = append(freqArr[count], num)
    }

    for i:=len(freqArr)-1; i >= 0; i-- {
        for _, n := range freqArr[i] {
            result = append(result, n)
            if len(result) == k {
                return result
            }
        }
    }

    return result
}

