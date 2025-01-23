# 4.3 Koko Eating Bananas `M`

[leetcode](https://leetcode.com/problems/koko-eating-bananas/) |
[youtube](https://www.youtube.com/watch?v=U2SozAs9RzA)

Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
The guards have gone and will come back in h hours.
Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile.
If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.
Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.
Return the minimum integer k such that she can eat all the bananas within h hours.

Example 1:
> - Input: piles = [3,6,7,11], h = 8
> - Output: 4

Example 2:
> - Input: piles = [30,11,23,4,20], h = 5
> - Output: 30

Example 3:
> - Input: piles = [30,11,23,4,20], h = 6
> - Output: 23

Constraints:
> - 1 <= piles.length <= 10^4
> - piles.length <= h <= 10^9
> - 1 <= piles[i] <= 10^9

<details>
  <summary><b>O(max*n) solution: brute force</b></summary>

- init max to max value of piles
- init res to max
- loop with k from 1 to max
  - init hours to 0
  - loop in piles incrementing hours by (current pile / k)
  - if hours is less or equal to h
    - set res to min of res and k
- return res
</details>

<details>
  <summary><b>O(log(max)*n) solution: binary search</b></summary>

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

```go
func minEatingSpeed(piles []int, h int) int {
    lo, hi := 1, slices.Max(piles)
    var res = hi

    for lo <= hi {
        mid := lo+(hi-lo)/2
        hours := 0 

        for _, pile := range piles {
            hours += int(math.Ceil(float64(pile)/float64(mid)))
        }

        if hours <= h {
            res = mid
            hi = mid-1
        } else {
            lo = mid+1
        }
    }

    return res
}
```

```ts
function minEatingSpeed(piles: number[], h: number): number {
    let lo = 1, hi = Math.max(...piles), res = hi

    while (lo <= hi) {
        const mid = lo + Math.floor((hi-lo)/2)
        let hours = 0
        for (const pile of piles) hours += Math.ceil(pile/mid)
        if (hours <= h) {
            res = mid
            hi = mid-1
        } else lo = mid+1
    }

    return res
};
```
</details>
