package sched

import (
	"sim/core"
)

type SRTF struct {
	name string
}

func (self SRTF) GetName() string {
	return self.name
}

func (self SRTF) Cmp(first, other core.Proc) bool {
	return first.Burst < other.Burst
}

func NewSRTF(name string) SRTF {
	return SRTF{
		name,
	}
}