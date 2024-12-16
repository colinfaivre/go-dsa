# 1.4 Group Anagrams `M`

[leetcode](https://leetcode.com/problems/group-anagrams/) |
[youtube](https://www.youtube.com/watch?v=vzdNOK2oB2E)

Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

> **Example 1:**
> - Input: strs = ["eat","tea","tan","ate","nat","bat"]
> - Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
>
> **Example 2:**
> - Input: strs = [""]
> - Output: [[""]]
>
> **Example 3:**
> - Input: strs = ["a"]
> - Output: [["a"]]
>
> **Constraints:**
> - 1 <= strs.length <= 10^4
> - 0 <= strs[i].length <= 100
> - strs[i] consists of lowercase English letters.

### hashmap solution: O(mn) (m: str avg length)
- init hashMap as:
  - key: arr 26 char counts (positioned by charcodes from 0->a to 26->z)
  - value: arr of strings containing anagrams
- init res as arr of arrs of strings containing anagrams
- loop in strs with value str
  - init charCount as an arr of 26 char counts
  - loop in str with value char
    - increment charCount at char - 'a'
  - append str to hashMap at charCount
- loop in hashmap with value anagrams
  - append anagrams to res
- return res

```go
func GroupAnagrams(strs []string) [][]string {
    hashMap := map[[26]int] []string {}
    res := [][]string {}

    for _, str := range strs {
        charCount := [26]int {}
        for _, char := range str {
            charCount[char - 'a']++
        }
        hashMap[charCount] = append(hashMap[charCount], str)
    }

    for _, anagrams := range hashMap {
        res = append(res, anagrams)
    }

    return res
}
```