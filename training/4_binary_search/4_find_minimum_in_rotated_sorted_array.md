# 4.4 Find Minimum In Rotated Sorted Array `M`

[leetcode](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) |
[youtube](https://www.youtube.com/watch?v=nIVW4P8b1VA)

Suppose an array of length n sorted in ascending order is rotated between 1 and n times.
For example, the array nums = [0,1,2,4,5,6,7] might become:
- [4,5,6,7,0,1,2] if it was rotated 4 times.
- [0,1,2,4,5,6,7] if it was rotated 7 times.

Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
Given the sorted rotated array nums of unique elements, return the minimum element of this array.
You must write an algorithm that runs in O(log n) time.

Example 1:
> - Input: nums = [3,4,5,1,2]
> - Output: 1
> - Explanation: The original array was [1,2,3,4,5] rotated 3 times.

Example 2:
> - Input: nums = [4,5,6,7,0,1,2]
> - Output: 0
> - Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

Example 3:
> - Input: nums = [11,13,15,17]
> - Output: 11
> - Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

Constraints:
> - n == nums.length
> - 1 <= n <= 5000
> - -5000 <= nums[i] <= 5000
> - All the integers of nums are unique.
> - nums is sorted and rotated between 1 and n times.

<details>
  <summary><b>O(n) Solution:</b></summary>

- init min to arbitrary value in nums
- loop in nums setting min to min of min and current value
- return min
</details>

<details>
  <summary><b>O(logn) Solution:</b></summary>

- init lo and hi to 0 and length of nums - 1
- init res to arbitrary value in nums
- loop while lo is less or equal to hi
  - if nums at lo is less than nums at hi
    - set res to min between min and nums at lo
    - break out of the loop
  - set mid to (lo+hi)/2
  - set res to min between res and nums at mid
  - if nums at mid is greater or equal to nums at hi
    - set lo to mid + 1
  - else set hi to mid - 1
- return res

```go
func FindMin(nums []int) int {
    lo, hi := 0, len(nums) - 1
    res := 50001

    for lo <= hi {
        if nums[lo] < nums[hi] {
            res = min(res, nums[lo])
            break
        }

        mid := (lo + hi) / 2
        res = min(res, nums[mid])
        
        if nums[mid] >= nums[hi] {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return res
}
```

```ts
function findMin(nums: number[]): number {
    let lo = 0, hi = nums.length-1, res = 50001
    while (lo <= hi) {
        if (nums[lo] < nums[hi]){
            res = nums[lo]
            break
        }
        const mid = lo + Math.floor((hi-lo)/2)
        res = Math.min(res, nums[mid])

        if (nums[mid] >= nums[hi]) lo = mid+1
        else hi = mid
    }
    return res
}
```
</details>
