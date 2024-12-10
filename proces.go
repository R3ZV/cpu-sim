package main

import (
	"container/heap"
	"fmt"
)

type proc struct {
	pid            int
	burst          int
	remainingBurst int

	// time until the process should execute again
	period   int
	priority int
}

func NewProc(pid int, burst int, period int) *proc {
	p := proc{
		pid:            pid,
		burst:          burst,
		remainingBurst: burst,
		period:         period,
		priority:       1,
	}

	return &p
}

type procHeap []proc

func (h procHeap) Len() int {
	return len(h)
}

func (h procHeap) Less(i, j int) bool {
	return h[i].burst < h[j].burst
}

func (h procHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *procHeap) Push(x interface{}) {
	*h = append(*h, x.(proc))
}

func (h *procHeap) Pop() interface{} {
	old := *h
	n := len(old)

	item := old[n-1]

	*h = old[0 : n-1]
	return item
}

func rateMonotonic(procese []proc) {
	// TODO:
	// MAX_TIME should be LCM
	const MAX_TIME = 200

	time := 0
	busyTime := 0

	q := &procHeap{}
    heap.Init(q)
    
	for i := 0; i < len(procese); i++ {
		heap.Push(q, procese[i])
	}

	for time < MAX_TIME {
		if q.Len() > 0 {
			busyTime++
			var currProces proc
			currProces = heap.Pop(q).(proc)

			currProces.remainingBurst--
			if currProces.remainingBurst != 0 {
				heap.Push(q, currProces)
			}
		}

		for i := 0; i < len(procese); i++ {
			if time%procese[i].period == 0 {
				heap.Push(q, procese[i])
			}
		}

		//TODO:
		// maybe someting more efficient?
		time++
	}
}

func main() {
	s := NewProc(1, 2, 3)
	fmt.Println(s.burst)
}
