# TODO

## 🏗️ Build
- ➡️ write a parser for `PrimsMST` test data
- add an implementation of `BinarySearchTree`
- quicksort with pivot selection or randomization

## 🧪 Test
- ➡️ improve error handling in `Heap` (methods should sometimes return errors)
- improve BFS and DFS tests with concrete problems to solve ?
- `Dijkstra()` should maybe take a copy of the graph instead of the pointer?
- improve and test error handling everywhere
- compare stack and queue slice based implementation with the LinkedList implementation (https://coderwall.com/p/cp5fya/measuring-execution-time-in-go)
- continue exploring `ginkgo`
- mocks
- BeforeEach()
- run just one file tests (one test suite per file ?)

## 📔 Document
- use https://go.dev/doc/comment on every file with:
  - applications of data structures and algo
  - `stanford dsa course` notes

## 🧗 Train
- ➡️ update problems/README.md
- ➡️ add two_sum sliding window solution
- add tests in `/problems` for sorting problems
- add copy_me template for scaffolding an algo from scratch to train
    - space complexity
    - time complexity
    - applications
    - implementations
- count inversions
- shortest paths with BFS
- SCCs on directed graph
- SCCs on undirected graph