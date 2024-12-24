# 5.2 Longest Substring Without Repeating Characters `M`

[leetcode](https://leetcode.com/problems/longest-substring-without-repeating-characters/) |
[youtube](https://www.youtube.com/watch?v=wiGpQwVHdE0)

Given a string s, find the length of the longest substring without repeating characters.

Example 1:
> - Input: s = "abcabcbb"
> - Output: 3
> - Explanation: The answer is "abc", with the length of 3.

Example 2:
> - Input: s = "bbbbb"
> - Output: 1
> - Explanation: The answer is "b", with the length of 1.

Example 3:
> - Input: s = "pwwkew"
> - Output: 3
> - Explanation: The answer is "wke", with the length of 3.
> Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:
> - 0 <= s.length <= 5 * 10^4
> - s consists of English letters, digits, symbols and spaces.

<details>
  <summary><b>O(n^2) brute-force solution:</b></summary>

- set res to 0
- loop in s with l
  - loop in s from l with r
    - if no repeating char
      - update res to max of res and r-l+1
- return res
</details>

<details>
  <summary><b>O(n) sliding window solution:</b></summary>

- init charSet to a set of chars
- init l to 0
- init res to 0
- loop in s with r
  - while s[r] is in set
    - remove s[l] from charSet
    - increment l
  - set charSet at s[r] to true
  - set res to max of res and r-l+1
- return res

```go
func LengthOfLongestSubstring(s string) int {
    charSet := make(map [byte]bool)
    l := 0
    res := 0

    for r := range s {
        for charSet[s[r]] {
            delete(charSet, s[l])
            l++
        }
        charSet[s[r]] = true
        res = max(res, r-l+1)
    }

    return res
}
```
</details>
