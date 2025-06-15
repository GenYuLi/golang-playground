package intheap

type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() (v any) {
	old := *h
	n := len(old)
	v = old[n-1]
	*h = old[0 : n-1]
	return v
}

// get the minimum element without removing it
func (h *IntMinHeap) Peek() int {
	if len(*h) == 0 {
		panic("heap is empty")
	}
	return (*h)[0]
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int {
	return len(h)
}

func (h IntMaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() (v any) {
	old := *h
	n := len(old)
	v = old[n-1]
	*h = old[0 : n-1]
	return v
}

// get the maximum element without removing it
func (h *IntMaxHeap) Peek() int {
	if len(*h) == 0 {
		panic("heap is empty")
	}
	return (*h)[0]
}
