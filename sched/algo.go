package sched

import (
	"sim/core"
)

type Scheduler interface {
	GetName() string
	Cmp(first, other core.Proc) bool
}
