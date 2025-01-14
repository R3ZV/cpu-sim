package sched

import (
	"sim/core"
)

type RM struct {
	Name string
}

func (self RM) GetName() string {
	return self.Name
}

func (self RM) IsPreemptive() bool {
	return true
}

func (self RM) IsRealTime() bool {
    return true
}

func (self RM) Cmp(first, second core.Proc, time int) bool {
	return first.Period < second.Period
}
func NewRM(name string) RM {
	return RM{
		name,
	}
}

func RMJobs(workload *[][]core.Proc) {
	jobs := []core.Proc{
		*core.NewProc(0, 35, 0, 0, 100),
		*core.NewProc(1, 20, 0, 0, 50),
		*core.NewProc(2, 20, 50, 0, 50),
		*core.NewProc(3, 20, 100, 0, 50),
		*core.NewProc(4, 35, 100, 0, 100),

	}
	*workload = append(*workload, jobs)

	jobs = []core.Proc{
		*core.NewProc(0, 25, 0, 0, 50),
		*core.NewProc(1, 35, 0, 0, 80),
		*core.NewProc(2, 25, 50, 0, 50),
		*core.NewProc(3, 35, 80, 0, 80),
		*core.NewProc(4, 25, 100, 0, 50),
		*core.NewProc(5, 25, 150, 0, 50),
		*core.NewProc(6, 35, 160, 0, 80),
		*core.NewProc(7, 25, 200, 0, 50),
		*core.NewProc(8, 35, 240, 0, 80),
		*core.NewProc(9, 25, 250, 0, 50),
		*core.NewProc(10, 35, 320, 0, 80),
		*core.NewProc(11, 35, 400, 0, 80),
	}

	*workload = append(*workload, jobs)
}
