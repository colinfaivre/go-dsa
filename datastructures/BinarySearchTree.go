// @MEDIUM https://ksw2000.medium.com/implement-a-binary-search-tree-with-go-969e716cc471
// @WIKI https://en.wikipedia.org/wiki/Binary_search_tree

package datastructures

// BSTNode represents the components of a binary search tree
type BSTNode struct {
	Left  *BSTNode
	Key   int
	Right *BSTNode
}

// BST represents a binary search tree
type BST struct {
	root *BSTNode
}

func (t *BST) GetRoot() *BSTNode {
	return t.root
}

// O(logn) Returns the minimum value of the tree
func (t *BST) Min() int {
	return t.root.minNode().Key
}
func (n BSTNode) minNode() BSTNode {
	for n.Left != nil {
		n = *n.Left
	}

	return n
}

// O(logn) Returns the minimum value of the tree
func (t *BST) Max() int {
	return t.root.maxNode().Key
}
func (n BSTNode) maxNode() BSTNode {
	for n.Right != nil {
		n = *n.Right
	}

	return n
}

// O(logn): Inserts a node to the tree
// the key to add should not be already in the tree
func (t *BST) Insert(k int) {
	if t.root != nil {
		t.root.insertNode(k)
	} else {
		t.root = &BSTNode{Left: nil, Right: nil, Key: k}
	}
}
func (n *BSTNode) insertNode(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &BSTNode{Key: k}
		} else {
			n.Right.insertNode(k)
		}
	} else {
		if n.Left == nil {
			n.Left = &BSTNode{Key: k}
		} else {
			n.Left.insertNode(k)
		}
	}
}

// O(logn): Searches for a key in the tree and returns true if a node has it
func (t BST) Search(k int) bool {
	return t.root.searchNode(k)
}
func (n *BSTNode) searchNode(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.searchNode(k)
	} else if n.Key > k {
		return n.Left.searchNode(k)
	}

	return true
}

func (t *BST) Remove(k int) {
	t.root.removeNode(k)
}

func (n *BSTNode) removeNode(k int) *BSTNode {
	if n == nil {
		return nil
	}

	if k < n.Key && n.Left != nil {
		n.Left = n.Left.removeNode(k)
		return n
	} else if k > n.Key && n.Right != nil {
		n.Right = n.Right.removeNode(k)
		return n
	} else {
		// remove a leaf node
		if n.Left == nil && n.Right == nil {
			n = nil
			return n
		}
		// remove a one-child node
		if n.Left == nil {
			n = n.Right
			return n
		} else if n.Right == nil {
			n = n.Left
			return n
		}
		// remove a two-children node
		aux := n.Right.minNode()
		n.Key = aux.Key
		n.Right = n.Right.removeNode(aux.Key)
		return n
	}
}

func (t *BST) InOrderTraverse(cb func(k int)) {
	t.root.inOrderTraverseNode(cb)
}

func (n *BSTNode) inOrderTraverseNode(cb func(k int)) {
	if n.Left != nil {
		n.Left.inOrderTraverseNode(cb)
	}

	cb(n.Key)

	if n.Right != nil {
		n.Right.inOrderTraverseNode(cb)
	}
}

func (t *BST) PreOrderTraverse(cb func(k int)) {
	t.root.preOrderTraverseNode(cb)
}

func (n *BSTNode) preOrderTraverseNode(cb func(k int)) {
	cb(n.Key)

	if n.Left != nil {
		n.Left.preOrderTraverseNode(cb)
	}

	if n.Right != nil {
		n.Right.preOrderTraverseNode(cb)
	}
}

func (t *BST) PostOrderTraverse(cb func(k int)) {
	t.root.postOrderTraverseNode(cb)
}

func (n *BSTNode) postOrderTraverseNode(cb func(k int)) {
	if n.Left != nil {
		n.Left.postOrderTraverseNode(cb)
	}

	if n.Right != nil {
		n.Right.postOrderTraverseNode(cb)
	}

	cb(n.Key)
}
