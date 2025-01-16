# 2.2 Two Sum II Input Array Is Sorted `M`

[leetcode](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) |
[youtube](https://www.youtube.com/watch?v=cQ1Oz4ckceM)

Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
The tests are generated such that there is exactly one solution. You may not use the same element twice.
Your solution must use only constant extra space.

Example 1:
> - Input: numbers = [2,7,11,15], target = 9
> - Output: [1,2]
> - Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Example 2:
> - Input: numbers = [2,3,4], target = 6
> - Output: [1,3]
> - Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Example 3:
> - Input: numbers = [-1,0], target = -1
> - Output: [1,2]
> - Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

Constraints:
> - 2 <= numbers.length <= 3 * 10^4
> - -1000 <= numbers[i] <= 1000
> - numbers is sorted in non-decreasing order.
> - -1000 <= target <= 1000
> - The tests are generated such that there is exactly one solution.

<details>
  <summary><b>O(n^2) solution - brute force</b></summary>

  - loop in nums
  - loop in nums
    - return the two pointers if sum equals target
</details>

<details>
  <summary><b>O(n) solution - 2 pointers</b></summary>

  - initiate left and right pointer at the first and last indexes of the array
  - loop while left is less than right
    - if the sum is greater than target decrement right
    - if the sum is less than target increment left
    - if the sum equals target return the two pointers
  
  ```go
  func TwoSum(numbers []int, target int) []int {
      l, r := 0, len(numbers) - 1
  
      for l < r {
          sum := numbers[l] + numbers[r]
          if sum < target { l++ }
          if sum > target { r-- }
          if sum == target { return []int{l+1, r+1} }
      }
  
      return []int{}
  }
  ```

```js
/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(numbers, target) {
    let lo = 0, hi = numbers.length - 1

    while (lo < hi) {
        const sum = numbers[lo] + numbers[hi]
        if (sum === target) return [lo + 1, hi + 1]
        sum > target ? hi-- : lo++
    }
}
```
</details>
