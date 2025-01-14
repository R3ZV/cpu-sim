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
	return first.Period-(time%first.Period) < other.Period-(time%other.Period)
}

func (self EDF) IsRealTime() bool {
	return true
}

func (self EDF) IsPreemptive() bool {
	return true
}
func NewEDF(name string) EDF {
	return EDF{
		name,
	}
}
