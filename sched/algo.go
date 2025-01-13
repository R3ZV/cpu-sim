package sched

import (
	"sim/core"
)

type Scheduler interface {
	GetName() string
	IsPreemptive() bool
	Cmp(first, other core.Proc, time int) bool
}
