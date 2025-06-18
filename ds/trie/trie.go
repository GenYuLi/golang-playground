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

// CollectWordsFrom a given node.
func (t *TrieNode) CollectWords(prefix string) []string {
	var words []string
	var dfs func(cur *TrieNode, currentWord string)

	dfs = func(cur *TrieNode, currentWord string) {
		if cur.IsWord {
			words = append(words, currentWord)
		}

		for ch, child := range cur.Children {
			dfs(child, currentWord+string(ch))
		}
	}

	// 1. 先走到 prefix 的結尾節點
	cur := t
	for _, ch := range prefix {
		if cur.Children[ch] == nil {
			return words // 找不到 prefix，直接回傳空 slice
		}
		cur = cur.Children[ch]
	}

	// 2. 從 prefix 結尾節點開始收集所有單字
	dfs(cur, prefix)
	return words
}
