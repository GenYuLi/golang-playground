package pq

import (
	"container/heap"

	"ds/pq/int_heap"
)

func Example() {
	h := &intheap.IntMaxHeap{2, 3, 7, 1, 4, 6, 5}
	heap.Init(h)
	println("Max Heap:", *h)
	println("Peek:", h.Peek())
	println("Pop:", heap.Pop(h).(int))

	mh := &intheap.IntMinHeap{2, 3, 7, 1, 4, 6, 5}
	heap.Init(mh)
	println("Min Heap:", *mh)
	println("Peek:", mh.Peek())
	println("Pop:", heap.Pop(mh).(int))
}
