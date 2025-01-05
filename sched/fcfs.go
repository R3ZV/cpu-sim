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

func (self FCFS) Cmp(first, other core.Proc) bool {
	return first.Priority < other.Priority
}

func NewFCFS(name string) FCFS {
	return FCFS{
		name,
	}
}
