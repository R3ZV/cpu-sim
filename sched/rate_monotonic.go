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
