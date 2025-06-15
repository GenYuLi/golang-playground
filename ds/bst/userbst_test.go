// userbst_test.go
package bst

import (
	"testing"
)

func dumpIDs(tree *BST[User]) []int64 {
	var ids []int64
	tree.InOrderTraversal(func(n *bst_node[User]) {
		ids = append(ids, n.Key.id)
	})
	return ids
}

func TestUserBST(t *testing.T) {
	ub := NewBST[User](Less)

	users := []User{
		{id: 42, name: "Alice"},
		{id: 7, name: "Bob"},
		{id: 99, name: "Carol"},
		{id: 7, name: "DupBob"}, // 重複 id
	}

	t.Run("Insert & Size", func(t *testing.T) {
		for i, u := range users[:3] {
			if ok := ub.Insert(u); !ok {
				t.Fatalf("Insert(%+v) #%d should succeed", u, i)
			}
		}
		if got, want := ub.Size(), 3; got != want {
			t.Fatalf("Size after 3 unique inserts = %d, want %d", got, want)
		}

		if ok := ub.Insert(users[3]); ok {
			t.Fatalf("Insert duplicate id should return false")
		}
		if got, want := ub.Size(), 3; got != want {
			t.Fatalf("Size after duplicate insert = %d, want %d", got, want)
		}
	})

	t.Run("Contains", func(t *testing.T) {
		if !ub.Contains(User{id: 42}) {
			t.Fatalf("Contains id=42 should be true")
		}
		if !ub.Contains(User{id: 7}) {
			t.Fatalf("Contains id=7 should be true")
		}
		if ub.Contains(User{id: 123}) {
			t.Fatalf("Contains id=123 should be false")
		}
	})

	t.Run("InOrderTraversal Order", func(t *testing.T) {
		got := dumpIDs(ub)
		want := []int64{7, 42, 99}
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("in-order index %d: got %d, want %d", i, got[i], want[i])
			}
		}
	})

	t.Run("Remove & SizeByTraverse", func(t *testing.T) {
		if ok := ub.Remove(User{id: 42}); !ok {
			t.Fatalf("Remove id=42 should succeed")
		}
		if got, want := ub.Size(), 2; got != want {
			t.Fatalf("Size after remove = %d, want %d", got, want)
		}
		if ub.Contains(User{id: 42}) {
			t.Fatalf("Contains id=42 after remove should be false")
		}

		if ok := ub.Remove(User{id: 999}); ok {
			t.Fatalf("Remove non-existent id should return false")
		}
		if got, want := ub.Size(), 2; got != want {
			t.Fatalf("Size after failed remove = %d, want %d", got, want)
		}

		if sbt := ub.SizeByTraverse(); sbt != ub.Size() {
			t.Fatalf("SizeByTraverse = %d, Size = %d; want equal", sbt, ub.Size())
		}
	})

	t.Run("Clear & IsEmpty", func(t *testing.T) {
		ub.Clear()
		if !ub.IsEmpty() {
			t.Fatalf("After Clear, IsEmpty should be true")
		}
		if ub.Size() != 0 {
			t.Fatalf("After Clear, Size = %d; want 0", ub.Size())
		}
	})
}

