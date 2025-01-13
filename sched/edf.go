package sched

import (
	"sim/core"
)

type EDF struct {
	name string
}

func (self EDF) GetName() string {
	return self.name
}

func (self EDF) Cmp(first, other core.Proc, time int) bool {
	return first.ParentPeriod-(time%first.ParentPeriod) < other.ParentPeriod-(time%other.ParentPeriod)
}
func (self EDF) IsPreemptive() bool {
	return true
}
func NewEDF(name string) EDF {
	return EDF{
		name,
	}
} //TODO give this man a test
