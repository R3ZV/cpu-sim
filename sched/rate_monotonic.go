package sched

import (
	"sim/core"
)

type RMS struct {
	Name string
}

func (self RMS) GetName() string {
	return self.Name
}
func (self RMS) IsPreemptive() bool {
	return true
}
func (self RMS) Cmp(first, second core.Proc, time int) bool {
	return first.ParentPeriod < second.ParentPeriod
}
func NewRMS(name string) RMS {
	return RMS{
		name,
	}
} //TODO give this man a test
