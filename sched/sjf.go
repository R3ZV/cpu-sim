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

func (self SJF) Cmp(first, other core.Proc) bool {
	return first.Burst < other.Burst
}

func NewSJF(name string) SJF {
	return SJF{
		name,
	}
}
