package sched

import (
	"sim/core"
)

type SJF struct {
	name string
}

func (self SJF) GetName() string {
	return self.name
}

func (self SJF) IsPreemptive() bool {
	return false
}

func (self SJF) Cmp(first, other core.Proc) bool {
	return first.Burst < other.Burst
}

func NewSJF(name string) SJF {
	return SJF{
		name,
	}
}

// Assumptions:
// 1. Same arrive time
// 2. Run to completion
// 3. No I/O
// 4. Known burst time
func SJFJobs(workload *[][]core.Proc) {
	jobs := []core.Proc{
		*core.NewProc(0, 3, 0, 0, -1),
		*core.NewProc(1, 1, 0, 0, -1),
		*core.NewProc(2, 4, 0, 0, -1),
	}
	*workload = append(*workload, jobs)

	jobs = []core.Proc{
		*core.NewProc(0, 30, 0, 0, -1),
		*core.NewProc(1, 30, 0, 0, -1),
		*core.NewProc(2, 100, 0, 0, -1),
	}

	*workload = append(*workload, jobs)
}
