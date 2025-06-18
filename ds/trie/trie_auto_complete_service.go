package trie

// coupling auto complete service with TrieNode
type AutocompleteService struct {
	dictionary *TrieNode // 直接依賴具體的 *TrieNode
}

func (s *AutocompleteService) Suggest(prefix string) []string {
	// ...
	return s.dictionary.CollectWords(prefix)
}

type DecouplingAutocompleteService struct {
	dictionary Trie // ✨ 依賴抽象的 interface，而不是具體的 struct
}

func (s *DecouplingAutocompleteService) Suggest(prefix string) []string {
	// ...
	return s.dictionary.CollectWords(prefix)
}
