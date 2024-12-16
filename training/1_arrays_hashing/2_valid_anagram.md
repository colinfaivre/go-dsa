# 1.2 Valid Anagram `E`

[leetcode](https://leetcode.com/problems/valid-anagram/)
[youtube](https://www.youtube.com/watch?v=9UtInBqnCgA)

Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

> **Example 1:**
> - Input: s = "anagram", t = "nagaram"
> - Output: true
>
> **Example 2:**
> - Input: s = "rat", t = "car"
> - Output: false
>
> **Constraints:**
> - 1 <= s.length, t.length <= 5 * 10^4
> - s and t consist of lowercase English letters.

### O(nlogn) - Space O(1): sort solution
- sort s
- sort t
- return s == t

### O(n) - Space O(n): hashmap solution
- return false if s and t have != length
- init countS, countT as hashmaps mapping chars to their count
- loop in s with i
  - increment countS at s[i]
  - increment countT at t[i]
- loop in countS with key
  - return false if countS at key != countT at key
- return true

```go
func IsAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    countS := map[byte] int {}
    countT := map[byte] int {}

    for i := range s {
        countS[s[i]]++
        countT[t[i]]++
    }

    for k := range countS {
        if countS[k] != countT[k] {
            return false
        }
    }

    return true
}
```
