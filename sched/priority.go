package sched

import (
	"sim/core"
)

type Priority struct {
	name string
}

func (self Priority) GetName() string {
	return self.name
}

func (self Priority) IsPreemptive() bool {
	return true
}

func (self Priority) Cmp(first, other core.Proc, time int) bool {
	return first.Priority < other.Priority
}

func NewPriority(name string) Priority {
	return Priority{
		name,
	}
}

func PriorityJobs(workload *[][]core.Proc) {
	jobs := []core.Proc{
		*core.NewProc(2, 100, 0, 1, -1),
		*core.NewProc(0, 50, 10, 3, -1),
		*core.NewProc(1, 10, 20, 2, -1),
	}
	*workload = append(*workload, jobs)

}
