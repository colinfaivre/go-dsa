package training

// 4.3 `M` Koko Eating Bananas

/*** @LEETCODE https://leetcode.com/problems/koko-eating-bananas/
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
The guards have gone and will come back in h hours.
Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile.
If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.
Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
Return the minimum integer k such that she can eat all the bananas within h hours.
---
Example 1:
Input: piles = [3,6,7,11], h = 8
Output: 4
---
Example 2:
Input: piles = [30,11,23,4,20], h = 5
Output: 30
---
Example 3:
Input: piles = [30,11,23,4,20], h = 6
Output: 23
---
Constraints:
1 <= piles.length <= 10^4
piles.length <= h <= 10^9
1 <= piles[i] <= 10^9
***/

/*** @SOLUTION https://www.youtube.com/watch?v=U2SozAs9RzA
O(max*n) solution: brute force
- init max to max value of piles
- init res to max
- loop with k from 1 to max
  - init hours to 0
  - loop in piles incrementing hours by (current pile / k)
  - if hours is less or equal to h
    - set res to min of res and k
- return res

O(log(max)*n) solution: binary search
- init left and right pointers to 1 and the max value of piles
- init res to right value
- loop while left is less or equal to right
  - init k (eating speed) to middle value from left to right
  - init hours to 0
  - loop in piles incrementing hours by (current pile / k)
  - if hours is less or equal to h
    - set res to min of res and k
    - set right pointer to k - 1
  - else set left pointer to k + 1
- return res
***/

func MinEatingSpeed(piles []int, h int) int {
    l, r := 1, slices.Max(piles)
    res := r

    for l <= r {
        k := (l + r) / 2
        hours := 0

        for _, p := range piles {
            hours += int(math.Ceil(float64(p) / float64(k)))
        }

        if hours <= h {
            res = min(res, k)
            r = k - 1
        } else {
            l = k + 1
        }
    }

    return res
}
