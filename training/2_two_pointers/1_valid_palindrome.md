# 2.1 Valid Palindrome `E`

[leetcode](https://leetcode.com/problems/valid-palindrome/) |
[youtube](https://www.youtube.com/watch?v=jJXJ16kPFWg)

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
> - Input: s = "A man, a plan, a canal: Panama"
> - Output: true
> - Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
> - Input: s = "race a car"
> - Output: false
> - Explanation: "raceacar" is not a palindrome.

Example 3:
> - Input: s = " "
> - Output: true
> - Explanation: s is an empty string "" after removing non-alphanumeric characters.
> - Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
> - 1 <= s.length <= 2 * 10^5
> - s consists only of printable ASCII characters.

<details>
<summary><b>O(n) solution</b></summary>
	
- set s to a lowercase alpha-numeric string  
- init l and r pointers to 0 and s length - 1
- loop while l < r
  - return false if s[l] != s[r]
  - increment l and decrement r
- return true
	
```go
func IsPalindrome(s string) bool {
    reg, _ := regexp.Compile("[^a-zA-Z0-9]+") 
    s = reg.ReplaceAllString(s, "")
    s = strings.ToLower(s)
    l, r := 0, len(s)-1

    for l < r {
	if s[l] != s[r] { return false }
	l++; r--
    }

    return true
}
```

```js
/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
    const cleanS = s.toLowerCase().replace(/[^0-9a-z]/gi, '').trim()

    let lo = 0, hi = cleanS.length - 1
    while (lo <= hi) {
        if (cleanS[lo] != cleanS[hi]) return false
        lo++
        hi--
    }

    return true
};
```
</details>
