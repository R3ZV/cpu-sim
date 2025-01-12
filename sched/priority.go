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

func (self Priority) Cmp(first, other core.Proc) bool {
	return first.Priority < other.Priority
}

func NewPriority(name string) Priority {
	return Priority{
		name,
	}
}
