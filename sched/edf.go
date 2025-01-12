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

func (self EDF) Cmp(first, other core.Proc) bool {
	return first.Priority < other.Priority
}

func NewEDF(name string) EDF {
	return EDF{
		name,
	}
}
