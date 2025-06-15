// bst_edge_test.go
package bst

import (
	"slices"
	"testing"
)

func dumpKeys(t *BST[int]) []int {
	var keys []int
	t.InOrderTraversal(func(n *bst_node[int]) { keys = append(keys, n.Key) })
	return keys
}

// --- 1. Edge cases for deleting root node  ------------------------------------

// (a) root only has one node (itself)
func TestRemoveRootSingleNode(t *testing.T) {
	tr := NewBST[int](func(a, b int) bool {
		return a < b
	})
	tr.Insert(42)

	if !tr.Remove(42) {
		t.Fatalf("failed to remove the only node (root)")
	}
	if !tr.IsEmpty() {
		t.Fatalf("tree should be empty after removing the only node")
	}
}

// (b) root only has left
func TestRemoveRootOnlyLeftChild(t *testing.T) {
	tr := NewBST[int](func(a, b int) bool {
		return a < b
	})
	tr.Insert(10)
	tr.Insert(5) // left child

	if !tr.Remove(10) {
		t.Fatalf("remove root with only left child failed")
	}
	if tr.Root.Key != 5 || tr.Size() != 1 {
		t.Fatalf("root should now be 5, size 1, got root=%v size=%d", tr.Root.Key, tr.Size())
	}
}

// (c) root only has right
func TestRemoveRootOnlyRightChild(t *testing.T) {
	tr := NewBST[int](func(a, b int) bool {
		return a < b
	})
	tr.Insert(10)
	tr.Insert(20)

	if !tr.Remove(10) {
		t.Fatalf("remove root with only right child failed")
	}
	if tr.Root.Key != 20 || tr.Size() != 1 {
		t.Fatalf("root should now be 20, size 1")
	}
}

// --- 2. Case where successor has a right child------------------------------------
//
// Tree：
//
//		     10
//			    \
//			   15
//			  /  \
//			12    18
//			  \
//			   13
//
//	 when delete 10, inorder-successor = 12, and 12 has right child 13
//	 After removal, 12 replaces 10, and 13 becomes right child of 12
func TestRemoveSuccessorHasRightChild(t *testing.T) {
	tr := NewBST[int](func(a, b int) bool {
		return a < b
	})
	for _, v := range []int{10, 15, 12, 18, 13} {
		tr.Insert(v)
	}
	if !tr.Remove(10) {
		t.Fatalf("failed to remove root 10")
	}

	want := []int{12, 13, 15, 18}
	got := dumpKeys(tr)
	if !slices.Equal(got, want) {
		t.Fatalf("in-order after removal incorrect: got %v, want %v", got, want)
	}
}

// --- 3. Pathological case: Decreasing insertions + Sequential deletion ---------------------------------

func TestSkewedInsertAndSequentialRemove(t *testing.T) {
	const n = 100
	tr := NewBST[int](func(a, b int) bool {
		return a < b
	})
	for i := n; i >= 1; i-- {
		tr.Insert(i)
	}
	if h := tr.Height(); h != n {
		t.Fatalf("left-skew height: got %d want %d", h, n)
	}
	for i := 1; i <= n; i++ {
		if !tr.Remove(i) {
			t.Fatalf("failed to remove %d in skewed tree", i)
		}
	}
	if !tr.IsEmpty() {
		t.Fatalf("tree not empty after removing all nodes in skewed test")
	}
}

// --- 4. repeated deletion of non-existent keys & reuse after clear()  ----------------------------

func TestRemoveNonExistAndReuse(t *testing.T) {
	tr := NewBST[int](func(a, b int) bool {
		return a < b
	})
	for _, v := range []int{3, 1, 4} {
		tr.Insert(v)
	}

	// try removing non-existent keys
	for i := 0; i < 5; i++ {
		if tr.Remove(99) {
			t.Fatalf("unexpected true when removing non-existent key")
		}
	}
	if tr.Size() != 3 {
		t.Fatalf("size should stay 3 after failed removals")
	}

	// Clear then rebuild
	tr.Clear()
	if !tr.IsEmpty() {
		t.Fatalf("Clear failed to empty the tree")
	}
	for _, v := range []int{8, 6, 9} {
		tr.Insert(v)
	}
	got := dumpKeys(tr)
	want := []int{6, 8, 9}
	if !slices.Equal(got, want) {
		t.Fatalf("tree rebuild mismatch, got %v want %v", got, want)
	}
}

// --- 5. Extreme values (minimum and maximum int) -------------------------------------

// int may be 32/64 bits, so we use bitwise to find min/max
func TestInsertMinMaxInt(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	tr := NewBST[int](func(a, b int) bool {
		return a < b
	})
	tr.Insert(0)
	tr.Insert(maxInt)
	tr.Insert(minInt)

	want := []int{minInt, 0, maxInt}
	got := dumpKeys(tr)
	if !slices.Equal(got, want) {
		t.Fatalf("failed min/max int ordering: got %v want %v", got, want)
	}

	// Try remove them one by one
	if !tr.Remove(maxInt) || !tr.Remove(minInt) || !tr.Remove(0) || !tr.IsEmpty() {
		t.Fatalf("remove min/max int sequence failed")
	}
}
