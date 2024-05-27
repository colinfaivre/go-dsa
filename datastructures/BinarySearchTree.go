// @MEDIUM https://ksw2000.medium.com/implement-a-binary-search-tree-with-go-969e716cc471
// @WIKI https://en.wikipedia.org/wiki/Binary_search_tree

package datastructures

// BSTNode represents the components of a binary search tree
type BSTNode struct {
	Left  *BSTNode
	Key   int
	Right *BSTNode
}

type BST struct {
	root *BSTNode
}

type BSTService interface {
	Insert(data int)
	Search(key int) bool
	GetRoot() *BSTNode
	InOrderTraversal() string
	PreOrderTraversal() string
	PostOrderTraversal() string
	Min()
	Max()
	Remove(key int)
}

func (t *BST) GetRoot() *BSTNode {
	return t.root
}

// O(logn) Returns the minimum value of the tree
func (t *BST) Min() int {
	return t.root.min()
}

func (n BSTNode) min() int {
	for n.Left != nil {
		n = *n.Left
	}

	return n.Key
}

// O(logn) Returns the minimum value of the tree
func (t *BST) Max() int {
	return t.root.max()
}

func (n BSTNode) max() int {
	for n.Right != nil {
		n = *n.Right
	}

	return n.Key
}

// O(logn): Inserts a node to the tree
// the key to add should not be already in the tree
func (t *BST) Insert(k int) {
	if t.root != nil {
		t.root.insert(k)
	} else {
		t.root = &BSTNode{Left: nil, Right: nil, Key: k}
	}
}

func (n *BSTNode) insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &BSTNode{Key: k}
		} else {
			n.Right.insert(k)
		}
	} else {
		if n.Left == nil {
			n.Left = &BSTNode{Key: k}
		} else {
			n.Left.insert(k)
		}
	}
}

// O(logn): Searches for a key in the tree and returns true if a node has it
func (t BST) Search(k int) bool {
	return t.root.search(k)
}

func (n *BSTNode) search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.search(k)
	} else if n.Key > k {
		return n.Left.search(k)
	}

	return true
}
