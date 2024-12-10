package main

import (
	"sim/core"
	"sim/sched"
)

func main() {
	procs := make([]core.Proc, 10)

	for i := 0; i < 10; i++ {
		procs[i] = core.Proc{
			Pid:            i,
			Burst:          10,
			RemainingBurst: 10,
			Period:         20,
			Priority:       1,
		}
	}

	sched.RateMonotonic(procs)
}
