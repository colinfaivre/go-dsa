# 5.3 Longest Repeating Character Replacement `M`

[leetcode](https://leetcode.com/problems/longest-repeating-character-replacement/) |
[youtube](https://www.youtube.com/watch?v=gqXU1UyA8pk)

You are given a string s and an integer k.
You can choose any character of the string and change it to any other uppercase English character.
You can perform this operation at most k times.
Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example 1:
> - Input: s = "ABAB", k = 2
> - Output: 4
> - Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:
> - Input: s = "AABABBA", k = 1
> - Output: 4
> - Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
> The substring "BBBB" has the longest repeating letters, which is 4.
> There may exists other ways to achieve this answer too.

Constraints:
> - 1 <= s.length <= 105
> - s consists of only uppercase English letters.
> - 0 <= k <= s.length

<details>
    <summary><b>O(26*n) Solution:</b></summary>

- init res at 0
- init l pointer at 0
- init charCount as an empty map from chars to integers
- loop in s w/ r pointer
    - if charCount at s[r] exists increment it
    - else init charCount at s[r] to 1
    - loop while window size minus max count from charMap is greater than k
        - decrement charCount at s[r]
        - increment l pointer
    - set res to max of res and window size
- return res

```go
func CharacterReplacement(s string, k int) int {
    res := 0
    l := 0
    charCount := map[byte]int{}

    for r := 0; r < len(s); r++ {
        if _, ok := charCount[s[r]]; ok {
            charCount[s[r]]++
        } else {
            charCount[s[r]] = 1
        }

        for (r - l + 1) - slices.Max(slices.Collect(maps.Values(charCount))) > k {
            charCount[s[l]]--
            l++
        }

        res = max(res, r - l + 1)
    }

    return res
}
```
</details>
