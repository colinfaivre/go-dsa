// @WIKI https://en.wikipedia.org/wiki/Trie
// @YOUTUBE https://www.youtube.com/watch?v=nL7BHR5vJDc

// @TODO Insert batch of words
// @TODO Delete word

package datastructures

const AlphabetSize = 26

// Node represents each node in the trie
type TrieNode struct {
	Children [AlphabetSize]*TrieNode
	IsEnd    bool
}

// Trie
type Trie struct {
	root *TrieNode
}

// Create a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (t *Trie) GetRoot() *TrieNode {
	return t.root
}

// O(w length): Insert takes a word and adds it into the Trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currNode.Children[charIndex] == nil {
			currNode.Children[charIndex] = &TrieNode{}
		}
		currNode = currNode.Children[charIndex]
	}

	currNode.IsEnd = true
}

// O(w length): Search takes a word and returns true if it is in the Trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currNode.Children[charIndex] == nil {
			return false
		}
		currNode = currNode.Children[charIndex]
	}

	return currNode.IsEnd
}
