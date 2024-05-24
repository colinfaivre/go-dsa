package datastructures

// BSTNode represents the components of a binary search tree
type BSTNode struct {
	left  *BSTNode
	key   int
	right *BSTNode
}

type BST struct {
	root *BSTNode
}

type BSTService interface {
	Insert(data int)
	Search(key int) bool
	GetRoot() *BSTNode
}

// Inserts adds a node to the tree
// the key to add should not be already in the tree
func (n *BSTNode) Insert(k int) {
	if n.key < k {
		if n.right == nil {
			n.right = &BSTNode{key: k}
		} else {
			n.right.Insert(k)
		}
	} else {
		if n.left == nil {
			n.left = &BSTNode{key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

// Searches for a key in the tree and returns true if a node has it
func (n *BSTNode) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.key < k {
		return n.right.Search(k)
	} else if n.key > k {
		return n.left.Search(k)
	}

	return true
}
