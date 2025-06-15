package bst

/*
Binary Search Tree (BST) implementation in Go.
this implementation only allows unique keys.
*/
type BST[T any] struct {
	Root  *bst_node[T]
	_size int // Optional: to keep track of the number of nodes
	less  func(a, b T) bool
}

func NewBST[T any](less func(a, b T) bool) *BST[T] {
	return &BST[T]{
		less: less,
	}
}

func (tree *BST[T]) Insert(key T) bool {
	var success bool
	if tree.Root == nil {
		tree.Root = &bst_node[T]{Key: key}
		success = true
	} else {
		success = tree.Root.insert(key, tree.less)
	}
	if success {
		tree._size++ // Increment size only on successful insert
	}
	return success
}

func (tree *BST[T]) Contains(key T) bool {
	if tree.Root == nil {
		return false
	}
	return tree.Root.search(key, tree.less) != nil
}

func (tree *BST[T]) Remove(key T) bool {
	if tree.Root == nil {
		return false
	}

	// call remove on internal root node function
	// and catch "new root" and whether deletion was successful
	newRoot, ok := tree.Root.remove(key, tree.less)
	if !ok {
		return false
	}
	tree.Root = newRoot
	tree._size-- // Decrement size only on successful removal
	return true
}

func (tree *BST[T]) Search(key T) *bst_node[T] {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.search(key, tree.less)
}

func (tree *BST[T]) InOrderTraversal(visit func(*bst_node[T])) {
	if tree.Root == nil {
		return
	}
	tree.in_order_traversal_helper(tree.Root, visit)
}

func (tree *BST[T]) IsEmpty() bool {
	return tree.Root == nil
}

func (tree *BST[T]) SizeByTraverse() int {
	if tree.Root == nil {
		return 0
	}
	var count int
	tree.InOrderTraversal(func(_ *bst_node[T]) {
		count++
	})
	return count
}

func (tree *BST[T]) Size() int {
	return tree._size
}

func (tree *BST[T]) Clear() {
	tree.Root = nil
	tree._size = 0 // Reset size to 0
}

func (tree *BST[T]) Height() int {
	if tree.Root == nil {
		return 0
	}
	return tree.rootHeight(tree.Root)
}

type bst_node[T any] struct {
	Key   T
	Left  *bst_node[T]
	Right *bst_node[T]
}

func (node *bst_node[T]) insert(key T, less func(a, b T) bool) bool {
	if less(key, node.Key) {
		if node.Left == nil {
			node.Left = &bst_node[T]{Key: key}
		} else {
			return node.Left.insert(key, less)
		}
	} else if less(node.Key, key) {
		if node.Right == nil {
			node.Right = &bst_node[T]{Key: key}
		} else {
			return node.Right.insert(key, less)
		}
	} else {
		// Key already exists, do not insert duplicates
		return false
	}
	return true
}

func (node *bst_node[T]) search(key T, less func(a, b T) bool) *bst_node[T] {
	if node == nil {
		return nil
	}
	if less(key, node.Key) {
		return node.Left.search(key, less)
	}
	if less(node.Key, key) {
		return node.Right.search(key, less)
	}
	return node // Key found, return the node
}

func (node *bst_node[T]) remove(key T, less func(a, b T) bool) (*bst_node[T], bool) {
	if node == nil {
		return nil, false // key not found
	}

	switch {
	case less(key, node.Key):
		newLeft, ok := node.Left.remove(key, less)
		node.Left = newLeft
		return node, ok

	case less(node.Key, key):
		newRight, ok := node.Right.remove(key, less)
		node.Right = newRight
		return node, ok

	default: // target found
		// 1. Only one side or no child →  return otherside or nil
		if node.Left == nil {
			return node.Right, true
		}
		if node.Right == nil {
			return node.Left, true
		}

		// 2. 2 child node →  find inorder successor (left minimum／ right maximum)
		successor := node.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		// overwrite by successor.Key, then remove successor recursively
		node.Key = successor.Key
		newRight, _ := node.Right.remove(successor.Key, less)
		node.Right = newRight
		return node, true
	}
}

func (tree *BST[T]) in_order_traversal_helper(node *bst_node[T], visit func(*bst_node[T])) {
	if node == nil {
		return
	}
	tree.in_order_traversal_helper(node.Left, visit)
	visit(node)
	tree.in_order_traversal_helper(node.Right, visit)
}

func (tree *BST[T]) rootHeight(node *bst_node[T]) int {
	if node == nil {
		return 0
	}
	leftHeight := tree.rootHeight(node.Left)
	rightHeight := tree.rootHeight(node.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
