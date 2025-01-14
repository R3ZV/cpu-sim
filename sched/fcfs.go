package sched

import (
	"sim/core"
)

type FCFS struct {
	name string
}

func (self FCFS) GetName() string {
	return self.name
}

func (self FCFS) IsPreemptive() bool {
	return false
}
func (self FCFS) IsRealTime() bool {
    return false
}

func (self FCFS) Cmp(first, other core.Proc, time int) bool {
	return first.Arrive < other.Arrive
}

func NewFCFS(name string) FCFS {
	return FCFS{
		name,
	}
}

// Assumptions:
// 1. Each job runs for the same amount of time.
// 2. Same arrive time
// 3. Run to completion
// 4. No I/O
// 5. Known burst time
func FCFSJobs(workload *[][]core.Proc) {
	jobs := []core.Proc{
		*core.NewProc(0, 10, 0, 0, -1),
		*core.NewProc(1, 10, 0, 0, -1),
		*core.NewProc(2, 10, 0, 0, -1),
	}

	*workload = append(*workload, jobs)

	// Relaxing assumption 1.
	jobs = []core.Proc{
		*core.NewProc(0, 100, 0, 0, -1),
		*core.NewProc(1, 10, 0, 0, -1),
		*core.NewProc(2, 10, 0, 0, -1),
	}
	*workload = append(*workload, jobs)
}
