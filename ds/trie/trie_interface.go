package trie

import "fmt"

// The reason why we need add an interface:
//  1. When we need different implementations of a Trie, we can define an interface.
//  2. To decrease coupling, we can use an interface to define the behavior of a Trie, and it
//     can provide ease of testing and flexibility in changing implementations.
//
// 宣告一個描述「Trie 應該有什麼行為」的 interface
type Trie interface {
	Insert(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
	CollectWords(prefix string) []string
}

// 這個函式不關心傳進來的是哪種 Trie，只要它會 Search 就行
func PrintSearchResult(trie Trie, word string) {
	if trie.Search(word) {
		fmt.Printf("'%s' found!\n", word)
	} else {
		fmt.Printf("'%s' not found.\n", word)
	}
}

func TestTrie() {
	originalTrie := NewTrie() // 回傳 *TrieNode，它實現了 Trie interface
	originalTrie.Insert("hello")

	// compressedTrie := NewCompressedTrie() // 假設這是你的新實作
	// compressedTrie.Insert("world")

	PrintSearchResult(originalTrie, "hello")
	// PrintSearchResult(compressedTrie, "world") // 也可以無痛換成另一種實作
}
