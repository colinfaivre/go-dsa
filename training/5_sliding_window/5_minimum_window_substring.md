# 5.5 Minimum Window Substring `H`

[leetcode](https://leetcode.com/problems/minimum-window-substring/) |
[youtube](https://www.youtube.com/watch?v=jSto0O4AJbM)

Given two strings s and t of lengths m and n respectively, return the minimum window substring
of s such that every character in t (including duplicates) is included in the window.
If there is no such substring, return the empty string "".
The testcases will be generated such that the answer is unique.

Example 1:
> - Input: s = "ADOBECODEBANC", t = "ABC"
> - Output: "BANC"
> - Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:
> - Input: s = "a", t = "a"
> - Output: "a"
> - Explanation: The entire string s is the minimum window.

Example 3:
> - Input: s = "a", t = "aa"
> - Output: ""
> - Explanation: Both 'a's from t must be included in the window.
> Since the largest window of s only has one 'a', return empty string.

Constraints:
> - m == s.length
> - n == t.length
> - 1 <= m, n <= 10^5
> - s and t consist of uppercase and lowercase English letters.

<details>
  <summary><b>O(n) solution - Sliding Window</b></summary>

```go
func minWindow(s string, t string) string {
    if len(s) == 0 || len(t) == 0 {
        return ""
    }

    // Create a map to store the frequency of characters in t
    targetFreq := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        targetFreq[t[i]]++
    }

    // Sliding window variables
    left, right := 0, 0
    minLen := len(s) + 1
    minStart := 0
    required := len(targetFreq)
    formed := 0

    // Current window frequency map
    windowFreq := make(map[byte]int)

    for right < len(s) {
        // Add the current character to the window
        char := s[right]
        windowFreq[char]++
        if targetFreq[char] > 0 && windowFreq[char] == targetFreq[char] {
            formed++
        }

        // Try to shrink the window
        for left <= right && formed == required {
            if right-left+1 < minLen {
                minLen = right - left + 1
                minStart = left
            }

            // Remove the character at the left
            leftChar := s[left]
            windowFreq[leftChar]--
            if targetFreq[leftChar] > 0 && windowFreq[leftChar] < targetFreq[leftChar] {
                formed--
            }
            left++
        }

        // Expand the window
        right++
    }

    if minLen == len(s)+1 {
        return ""
    }
    return s[minStart : minStart+minLen]
}
```
</details>
