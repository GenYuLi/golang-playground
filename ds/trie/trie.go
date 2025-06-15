// trie.go
package trie

// TrieNode represents one node of a Trie.
type TrieNode struct {
	Children map[rune]*TrieNode // outgoing edges
	IsWord   bool               // marks end-of-word
}

// NewTrie returns an empty root node.
func NewTrie() *TrieNode {
	return &TrieNode{Children: make(map[rune]*TrieNode)}
}

// Insert inserts word into the trie.
func (t *TrieNode) Insert(word string) {
	cur := t
	for _, ch := range word {
		if cur.Children[ch] == nil {
			cur.Children[ch] = NewTrie()
		}
		cur = cur.Children[ch]
	}
	cur.IsWord = true
}

// Search returns true if word exists in the trie.
func (t *TrieNode) Search(word string) bool {
	cur := t
	for _, ch := range word {
		cur = cur.Children[ch]
		if cur == nil {
			return false
		}
	}
	return cur.IsWord
}

// StartsWith returns true if any word in the trie has the given prefix.
func (t *TrieNode) StartsWith(prefix string) bool {
	cur := t
	for _, ch := range prefix {
		cur = cur.Children[ch]
		if cur == nil {
			return false
		}
	}
	return true
}
