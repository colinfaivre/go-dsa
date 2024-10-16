package training

// 5.4 `M` Permutation In String

/*** @LEETCODE https://leetcode.com/problems/permutation-in-string/
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
In other words, return true if one of s1's permutations is the substring of s2.
---
Example 1:
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
---
Example 2:
Input: s1 = "ab", s2 = "eidboaoo"
Output: false
---
Constraints:
1 <= s1.length, s2.length <= 10^4
s1 and s2 consist of lowercase English letters.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=UbyhOgBN834
O(26n) solution:
- return false if s1 is bigger than s2
- init cnt1 and cnt2 as two 26 sized arrays mapping alphabet letters to their count
- loop through s1 to fill cnt1 with its char count and cnt2 with s2 subset of s1 length char count
- return true if cnt1 == cnt2
- loop from s1 length to s2 length with i
  - increment cnt2 at s2[i]
  - decrement cnt2 at s2[i- s1 length]
  - return true if cnt1 == cnt2
- return false
***/

func CheckInclusion(s1 string, s2 string) bool {
    	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	
	var cnt1, cnt2 [26]int
	
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	
	for i := l1; i < l2; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-l1]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	
	return false
}
