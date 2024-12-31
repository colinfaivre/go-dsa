# 3.6 Car Fleet `M`

[leetcode](https://leetcode.com/problems/car-fleet/)
[youtube](https://www.youtube.com/watch?v=Pr6T-3yB9RM)

There are n cars at given miles away from the starting mile 0, traveling to reach the mile target.
You are given two integer array position and speed, both of length n, where position[i] is the starting mile of the ith car and speed[i] is the speed of the ith car in miles per hour.
A car cannot pass another car, but it can catch up and then travel next to it at the speed of the slower car.
A car fleet is a car or cars driving next to each other. The speed of the car fleet is the minimum speed of any car in the fleet.
If a car catches up to a car fleet at the mile target, it will still be considered as part of the car fleet.
Return the number of car fleets that will arrive at the destination.

Example 1:
> - Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
> - Output: 3
> - Explanation:
> 	- The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12. The fleet forms at target.
> 	- The car starting at 0 (speed 1) does not catch up to any other car, so it is a fleet by itself.
> 	- The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.

Example 2:
> - Input: target = 10, position = [3], speed = [3]
> - Output: 1
> - Explanation:
> 	- There is only one car, hence there is only one fleet.

Example 3:
> - Input: target = 100, position = [0,2,4], speed = [4,2,1]
> - Output: 1
> - Explanation:
> 	- The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4. The car starting at 4 (speed 1) travels to 5.
> 	- Then, the fleet at 4 (speed 2) and the car at position 5 (speed 1) become one fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.

Constraints:
> - n == position.length == speed.length
> - 1 <= n <= 10^5
> - 0 < target <= 10^6
> - 0 <= position[i] < target
> - All the values of position are unique.
> - 0 < speed[i] <= 10^6

<details>
	<summary><b>O(nlogn) solution - sort and stack</b></summary>

- store (position, speed) pairs in an array
- sort posSpeedPairs by increasing order of positions
- init a stack to store time to destination
- loop in posSpeedPairs in reverse order
	- push current car time to destination in stack (target-position)/speed
	- if current car time to destination is less than previous
		- pop current from the stack
- return stack length

```go
func CarFleet(target int, position []int, speed []int) int {
	posSpeedPairs := make([][2]int, len(position))
	for i := 0; i < len(position); i++ {
		posSpeedPairs[i] = [2]int{position[i], speed[i]}
	}
	sort.Slice(posSpeedPairs, func(i, j int) bool {
		return posSpeedPairs[i][0] < posSpeedPairs[j][0]
	})

	stack := []float64{} // Track the times to reach the target for each car

	for i := len(posSpeedPairs) - 1; i >= 0; i-- { // Traverse cars from last to first
		position := float64(posSpeedPairs[i][0])
		speed := float64(posSpeedPairs[i][1])
		timeToReach := (float64(target) - position) / speed // Time to reach the target for the current car
		
		if len(stack) == 0 || timeToReach > stack[len(stack)-1] {
			stack = append(stack, timeToReach)
		}
	}

	return len(stack)
}
```
</details>
