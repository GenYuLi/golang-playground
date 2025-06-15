// bst_test.go
package bst

import (
	"slices"
	"testing"
)

// helper: build tree from slice
func buildBST(nums []int) *BST[int] {
	t := NewBST[int](func(a, b int) bool {
		return a < b
	})
	for _, v := range nums {
		t.Insert(v)
	}
	return t
}

func TestInsertAndContains(t *testing.T) {
	tree := buildBST([]int{5, 2, 8, 1, 4, 7, 9})

	for _, v := range []int{1, 2, 4, 5, 7, 8, 9} {
		if !tree.Contains(v) {
			t.Fatalf("expected tree to contain %d", v)
		}
	}

	if tree.Contains(42) {
		t.Fatalf("tree should not contain 42")
	}
}

func TestDuplicateInsertRejected(t *testing.T) {
	tree := NewBST[int](func(a, b int) bool {
		return a < b
	})
	if !tree.Insert(10) {
		t.Fatalf("first insert must succeed")
	}
	if tree.Size() != 1 {
		t.Fatalf("size should be 1 after first insert")
	}

	if ok := tree.Insert(10); ok {
		t.Fatalf("duplicate insert should have returned false")
	}
	if tree.Size() != 1 {
		t.Fatalf("size should remain 1 after duplicate insert")
	}
}

func TestInOrderTraversal(t *testing.T) {
	data := []int{5, 2, 8, 1, 4, 7, 9}
	tree := buildBST(data)

	var got []int
	tree.InOrderTraversal(func(n *bst_node[int]) {
		got = append(got, n.Key)
	})

	want := slices.Clone(data)
	slices.Sort(want)

	if !slices.Equal(got, want) {
		t.Fatalf("in-order traversal mismatch, got %v, want %v", got, want)
	}
}

func TestRemoveCases(t *testing.T) {
	tree := buildBST([]int{5, 2, 8, 1, 4, 7, 9}) // balanced-ish

	// 1. remove leaf (1)
	if !tree.Remove(1) {
		t.Fatalf("failed to remove leaf 1")
	}
	if tree.Contains(1) || tree.Size() != 6 {
		t.Fatalf("leaf removal inconsistent")
	}

	// 2. remove node with one child (4)
	if !tree.Remove(4) {
		t.Fatalf("failed to remove node 4 with one child (nil right)")
	}
	if tree.Contains(4) || tree.Size() != 5 {
		t.Fatalf("single-child removal inconsistent")
	}

	// 3. remove node with two children (8)
	if !tree.Remove(8) {
		t.Fatalf("failed to remove node 8 with two children")
	}
	if tree.Contains(8) || tree.Size() != 4 {
		t.Fatalf("two-child removal inconsistent")
	}

	// 4. remove root (5) which now has two children (2,7 or 2,9 depending)
	if !tree.Remove(5) {
		t.Fatalf("failed to remove root")
	}
	if tree.Contains(5) || tree.Size() != 3 {
		t.Fatalf("root removal inconsistent")
	}

	// 5. remove non-existent key
	if tree.Remove(42) {
		t.Fatalf("removing 42 should return false")
	}
}

func TestHeightAndSize(t *testing.T) {
	tree := NewBST[int](func(a, b int) bool {
		return a < b
	})
	if h := tree.Height(); h != 0 {
		t.Fatalf("empty tree height should be 0")
	}
	if s := tree.Size(); s != 0 {
		t.Fatalf("empty tree size should be 0")
	}

	tree = buildBST([]int{10, 5, 15, 3, 7})
	if h := tree.Height(); h != 3 { // 10 -> 5 -> 3 is the longest path
		t.Fatalf("expected height 3, got %d", h)
	}
	if s := tree.Size(); s != 5 {
		t.Fatalf("expected size 5, got %d", s)
	}
}

func TestClearAndIsEmpty(t *testing.T) {
	tree := buildBST([]int{1, 2, 3})
	if tree.IsEmpty() {
		t.Fatalf("tree should not be empty")
	}
	tree.Clear()
	if !tree.IsEmpty() {
		t.Fatalf("tree should be empty after Clear")
	}
}

func TestGenericWithString(t *testing.T) {
	strs := []string{"delta", "alpha", "charlie", "bravo"}
	tree := NewBST[string](func(a, b string) bool {
		return a < b
	})
	for _, s := range strs {
		tree.Insert(s)
	}

	for _, s := range strs {
		if !tree.Contains(s) {
			t.Fatalf("tree should contain %s", s)
		}
	}

	var got []string
	tree.InOrderTraversal(func(n *bst_node[string]) { got = append(got, n.Key) })
	want := slices.Clone(strs)
	slices.SortFunc[[]string](want, func(a, b string) int { // Go1.22: SortFunc generic
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	if !slices.Equal(got, want) {
		t.Fatalf("string in-order mismatch: got %v, want %v", got, want)
	}
}
