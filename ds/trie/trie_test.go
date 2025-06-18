// service_test.go
package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockTrie struct{}

// 為了滿足 service 中 dictionary field 所依賴的 interface，
// 我們必須實作這個 interface 要求的所有方法。
// 根據上個問題，AutocompleteService 用到了 CollectWords。
func (m *MockTrie) CollectWords(prefix string) []string {
	if prefix == "go" {
		return []string{"golang", "gopher"}
	}
	return nil
}

// 為了讓 MockTrie 完整實現我們之前定義的 Trie interface，
// 把其他方法也補上空實作。
func (m *MockTrie) Insert(word string)            {}
func (m *MockTrie) Search(word string) bool       { return false }
func (m *MockTrie) StartsWith(prefix string) bool { return false }

func TestAutocompleteService(t *testing.T) {
	mock := &MockTrie{}
	service := &DecouplingAutocompleteService{dictionary: mock}

	suggestions := service.Suggest("go")
	// ... 進行斷言

	expected := []string{"golang", "gopher"}
	assert.Equal(t, expected, suggestions, "給定 'go' 前綴時，應該要回傳 'golang' 和 'gopher'")
	assert.Len(t, suggestions, 2, "建議單字的數量應該是 2")
}
