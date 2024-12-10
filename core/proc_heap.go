package core

type ProcHeap []Proc

func (h ProcHeap) Len() int {
	return len(h)
}

func (h ProcHeap) Less(i, j int) bool {
	return h[i].Burst < h[j].Burst
}

func (h ProcHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ProcHeap) Push(x interface{}) {
	*h = append(*h, x.(Proc))
}

func (h *ProcHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]

	*h = old[0 : n-1]
	return item
}
