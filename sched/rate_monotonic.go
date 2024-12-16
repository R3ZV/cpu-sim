package sched


 import (
 	"container/heap"
 )

 // Rate monotonic uses preemption.
 // It assumes that every time a process aquires the CPU
 // the burst is always the same.
 //
 // We can test whether a set of processes can be completed
 // until the deadline by computing the sum of burst_i / period_i
 func gcd(a, b int){
	if (b == 0){
	   return a
	}
	return gcd(b, a%b)
 }
 func LCM(x, y int){
	return x*y/gcd(x, y)
 }
 func commonProcessCycle(procs []core.metaProc){
     leastCycle :=1
	 for i :=0; i<len(procs); i++{
		leastCycle=lcm(leastCycle, procs[i].Period)
	 }
	 return leastCycle
 }
 func RateMonotonic(procs []core.Proc) {
	//sanity check: if the processes take too long overall, we don't even bother to arrange them
	totalRatio := 0
	for i:=0; i<len(procs); i++{
		totalRatio += procs[i].burst/procs[i].period;
	} 
	if(totalRatio>1){//if we can't schedule all processes
        panic("We can't execute all these processes in due time")
	}
 	const MAX_TIME = commonProcessCycle(procs)
    
 	time := 0
 	busyTime := 0

 	q := &core.ProcHeap{}
 	heap.Init(q)

 	for i := 0; i < len(procs); i++ {
 		heap.Push(q, procs[i])
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

 		for i := 0; i < len(procs); i++ {
 			if time%procs[i].Period == 0 {
 				heap.Push(q, procs[i])
 			}
 		}
 		//TODO:
 		 //maybe someting more efficient?
 		time++
 	}
 }
