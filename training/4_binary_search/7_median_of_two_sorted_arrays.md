# 4.7 Median of Two Sorted Arrays `H`

[leetcode](https://leetcode.com/problems/median-of-two-sorted-arrays/) |
[youtube](https://www.youtube.com/watch?v=q6IEA26hvXc)

Given two sorted arrays nums1 and nums2 of size m and n respectively,
return the median of the two sorted arrays.
The overall run time complexity should be O(log (m+n)).

Example 1:
> - Input: nums1 = [1,3], nums2 = [2]
> - Output: 2.00000
> - Explanation: merged array = [1,2,3] and median is 2.

Example 2:
> - Input: nums1 = [1,2], nums2 = [3,4]
> - Output: 2.50000
> - Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

Constraints:
> - nums1.length == m
> - nums2.length == n
> - 0 <= m <= 1000
> - 0 <= n <= 1000
> - 1 <= m + n <= 2000
> - -10^6 <= nums1[i], nums2[i] <= 10^6

<details>
  <summary><b>O(log(m+n)) solution - binary search</b></summary>

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) { return findMedianSortedArrays(nums2, nums1) }
    m, n := len(nums1), len(nums2)
    left, right := 0, m
    halfLen := (m + n + 1) / 2

    for left <= right {
        i := (left + right) / 2
        j := halfLen - i

        if i < m && nums1[i] < nums2[j-1] {
            left = i + 1 // i is too small, increase it
        } else if i > 0 && nums1[i-1] > nums2[j] {
            right = i - 1 // i is too big, decrease it
        } else { // i is perfect
            var maxLeft int
            if i == 0 {
                maxLeft = nums2[j-1]
            } else if j == 0 {
                maxLeft = nums1[i-1]
            } else {
                maxLeft = max(nums1[i-1], nums2[j-1])
            }

            if (m+n)%2 == 1 { return float64(maxLeft) }

            var minRight int
            if i == m {
                minRight = nums2[j]
            } else if j == n {
                minRight = nums1[i]
            } else {
                minRight = min(nums1[i], nums2[j])
            }

            return float64(maxLeft+minRight) / 2.0
        }
    }

    return 0.0
}
```
</details>
