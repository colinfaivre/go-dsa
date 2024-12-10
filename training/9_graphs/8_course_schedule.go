package training

// 9.8 `M` Course Schedule

/*** @LEETCODE https://leetcode.com/problems/course-schedule/
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
You are given an array prerequisites where
prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.
For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.
---
Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.
---
Example 2:
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
---
Constraints:
1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
All the pairs prerequisites[i] are unique.
***/

/*** @SOLUTION https://www.youtube.com/watch?v=EgI5nU9etnU
***/

func canFinish(numCourses int, prerequisites [][]int) bool {
    adjList := make(map[int] []int)
    for _, pre := range prerequisites {
        if _, exists := adjList[pre[0]]; !exists {
            adjList[pre[0]] = []int{ pre[1] }
        } else {
            adjList[pre[0]] = append(adjList[pre[0]], pre[1])
        }
    }
    visited := make(map[int] bool)
    for c := 0; c < numCourses; c++ {
        if !noCycle(c, visited, adjList) {
            return false
        }
    }

    return true
}

func noCycle(course int, visited map[int] bool, adjList map[int] []int) bool {
    if visited[course] { return false }
    if len(adjList[course]) == 0 { return true }
    visited[course] = true
    for _, pre := range adjList[course] {
        if !noCycle(pre, visited, adjList) { return false }
    }
    delete(visited, course)
    adjList[course] = []int{}
    return true
}
