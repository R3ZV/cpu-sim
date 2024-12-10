package sched

import (
	"sim/core"

	"container/heap"
)

func RateMonotonic(procese []core.Proc) {
	// TODO:
	// MAX_TIME should be LCM
	const MAX_TIME = 200

	time := 0
	busyTime := 0

	q := &core.ProcHeap{}
	heap.Init(q)

	for i := 0; i < len(procese); i++ {
		heap.Push(q, procese[i])
	}

	for time < MAX_TIME {
		if q.Len() > 0 {
			busyTime++
			var currProces core.Proc
			currProces = heap.Pop(q).(core.Proc)

			currProces.RemainingBurst--
			if currProces.RemainingBurst != 0 {
				heap.Push(q, currProces)
			}
		}

		for i := 0; i < len(procese); i++ {
			if time%procese[i].Period == 0 {
				heap.Push(q, procese[i])
			}
		}

		//TODO:
		// maybe someting more efficient?
		time++
	}
}
