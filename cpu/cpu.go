package cpu

import (
	"sim/core"
	"sim/sched"

	"container/heap"
	"fmt"
)

type CPU struct {
	timer       int
	activeTimer int
	currProc    *core.Proc
	procs       *core.ProcHeap
}

func NewCPU(schedAlgo sched.Scheduler) *CPU {
	procHeap := core.NewProcHeap(schedAlgo.Cmp)
	heap.Init(procHeap)

	return &CPU{
		timer:       0,
		activeTimer: 0,
		currProc:    nil,
		procs:       procHeap,
	}
}

func (self *CPU) loadProc() {
	self.currProc = heap.Pop(self.procs).(*core.Proc)
	fmt.Printf("Loading proc PID: %d\n", self.currProc.Pid)
}

func (self *CPU) unloadProc() {
	fmt.Printf("Unloading proc PID: %d\n", self.currProc.Pid)
	self.currProc = nil
}

func (self *CPU) AddProc(proc core.Proc) {
	heap.Push(self.procs, proc)
}

func (self *CPU) available() bool {
	return self.currProc == nil
}

func (self *CPU) procDone() bool {
	if self.currProc == nil {
		panic("ProcDone should only be called if there was a process on CPU")
	}
	return self.currProc.Burst == 0
}

func (self *CPU) Tick() {
	self.timer++
	if self.available() {
		self.loadProc()
	}

	if self.procDone() {
		self.unloadProc()
	}

	if self.currProc != nil {
		self.currProc.Burst--
	}
	self.activeTimer++

}

func (self *CPU) IsDone() bool {
	return self.procs.Len() == 0 && self.currProc == nil
}

func (self *CPU) Usage() float32 {
	return (float32(self.activeTimer) / float32(self.timer)) * 100.0
}
