package sched

import (
	"sim/core"
)

type PSJF struct {
	name string
}

func (self PSJF) GetName() string {
	return self.name
}

func (self PSJF) IsPreemptive() bool {
	return true
}

func (self PSJF) Cmp(first, other core.Proc, time int) bool {
	return first.Burst < other.Burst
}

func NewPSJF(name string) PSJF {
	return PSJF{
		name,
	}
}

// Assumptions:
// 1. Run to completion
// 2. No I/O
// 3. Known burst time
func PSJFJobs(workload *[][]core.Proc) {
	// Relaxing assumption 1.
	jobs := []core.Proc{
		*core.NewProc(2, 100, 0, 0, -1),
		*core.NewProc(0, 50, 10, 0, -1),
		*core.NewProc(1, 10, 20, 0, -1),
	}

	*workload = append(*workload, jobs)
}
