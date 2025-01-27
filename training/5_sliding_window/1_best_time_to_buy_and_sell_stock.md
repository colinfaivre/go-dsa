# 5.1 Best Time to Buy And Sell Stock `E`

[leetcode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) |
[youtube](https://www.youtube.com/watch?v=1pkOgXD63yU)

You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:
> - Input: prices = [7,1,5,3,6,4]
> - Output: 5
> - Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
> Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
> - Input: prices = [7,6,4,3,1]
> - Output: 0
> - Explanation: In this case, no transactions are done and the max profit = 0.

Constraints:
> - 1 <= prices.length <= 10^5
> - 0 <= prices[i] <= 10^4

<details>
  <summary><b>O(n): sliding window</b></summary>

- init left and right pointers at 0 and 1
- init max profit at 0
- loop in prices while right pointes is less than profits length
  - if the transaction is profitable (prices[l] < prices[r])
    - update max profit
  - else set left pointer to right pointer value
  - increment right pointer

```go
func MaxProfit(prices []int) int {
    l, r := 0, 1
    maxP := 0

    for r < len(prices) {
        if prices[l] < prices[r] {
            profit := prices[r] - prices[l]
            maxP = max(maxP, profit)
        } else {
            l = r
        }
        r++
    }

    return maxP
}
```

```ts
function maxProfit(prices: number[]): number {
    let lo = 0, hi = 0, maxProfit = 0
    while (hi < prices.length) {
        if (prices[lo] < prices[hi]) maxProfit = Math.max(maxProfit, prices[hi]-prices[lo])
        else lo = hi
        hi++
    }

    return maxProfit
}
```
</details>
