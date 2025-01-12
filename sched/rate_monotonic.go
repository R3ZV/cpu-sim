package sched

// import (
// 	"container/heap"
// )

// / Rate monotonic uses preemption.
// / It assumes that every time a process aquires the CPU
// / the burst is always the same.
// /
// / We can test weather a set of processes can be completed
// / until the deadline by computing the sum of burst_i / period_i
// func RateMonotonic(procs []core.Proc) {
// 	// TODO:
// 	// MAX_TIME should be LCM
// 	const MAX_TIME = 200
//
// 	time := 0
// 	busyTime := 0
//
// 	q := &core.ProcHeap{}
// 	heap.Init(q)
//
// 	for i := 0; i < len(procs); i++ {
// 		heap.Push(q, procs[i])
// 	}
//
// 	for time < MAX_TIME {
// 		if q.Len() > 0 {
// 			busyTime++
// 			var currProces core.Proc
// 			currProces = heap.Pop(q).(core.Proc)
//
// 			currProces.RemainingBurst--
// 			if currProces.RemainingBurst != 0 {
// 				heap.Push(q, currProces)
// 			}
// 		}
//
// 		for i := 0; i < len(procs); i++ {
// 			if time%procs[i].Period == 0 {
// 				heap.Push(q, procs[i])
// 			}
// 		}
//
// 		//TODO:
// 		// maybe someting more efficient?
// 		time++
// 	}
// }
