package cpu

import (
	"sim/core"
	"sim/log"
	"sim/sched"

	"container/heap"
)

type CPU struct {
	timer        int
	activeTimer  int
	currProc     *core.Proc
	Procs        *core.ProcHeap
	taTimes      []int
	waitTimes    []int
	arrivalTimes []int
}

// auxiliary function to check if sonething is an arrival time
func exists(l []int, key int) bool {
	for i := range l {
		if l[i] == key {
			return true
		}
	}
	return false
}

func NewCPU(schedAlgo sched.Scheduler) *CPU {
	procHeap := core.NewProcHeap(schedAlgo.Cmp)
	heap.Init(procHeap)

	return &CPU{
		timer:       0,
		activeTimer: 0,
		currProc:    nil,
		Procs:       procHeap,
	}
}

func (self *CPU) loadProc() {
	self.currProc = heap.Pop(self.Procs).(*core.Proc)
	log.Debug("Loading proc PID: %d\n", self.currProc.Pid)
}

func (self *CPU) unloadProc() {
	log.Debug("Unloading proc PID: %d\n", self.currProc.Pid)
	self.currProc = nil
}

func (self *CPU) changeProc() {
	log.Debug("Unloading proc PID: %d\n", self.currProc.Pid)
	exchange := heap.Pop(self.Procs).(*core.Proc)
	log.Debug("Loading proc PID: %d\n", exchange.Pid)
	heap.Push(self.Procs, *self.currProc)
	self.currProc = exchange
}

func (self *CPU) AddProc(proc core.Proc) {
	heap.Push(self.Procs, proc)
	self.arrivalTimes = append(self.arrivalTimes, proc.Arrive)
}

func (self *CPU) available() bool {
	return self.currProc == nil
}

func (self *CPU) procDone() bool {
    log.Assert(self.currProc != nil, "ProcDone should only be called if there was a process on CPU")
	return self.currProc.Burst == 0
}

func (self *CPU) Tick() {
	if self.available() {
		self.loadProc()
		self.waitTimes = append(self.waitTimes, self.timer-self.currProc.Arrive)
	}

	if self.procDone() {
		self.taTimes = append(self.taTimes, self.timer-self.currProc.Arrive)
		self.unloadProc()
		self.timer--
	}

	if self.currProc != nil {
		self.currProc.Burst--
		self.activeTimer++
	}
	self.timer++
}

// we need a different Tick() to implement preemption,
// since we can pop a process that isn't done, the above can't support it
func (self *CPU) PreemptiveTick() {
	if self.available() {
		self.loadProc()
	}

	if self.procDone() {
		self.taTimes = append(self.taTimes, self.timer-self.currProc.Arrive)
		self.waitTimes = append(self.waitTimes, self.timer-self.currProc.Arrive-self.currProc.InitBurst)
		self.unloadProc()
		self.timer--
	}

	// if exists(self.arrivalTimes, self.timer) {
	// 	self.changeProc()
	// }

	if self.currProc != nil {
		self.currProc.Burst--
		self.activeTimer++
	}
	self.timer++
}

func (self *CPU) IsDone() bool {
	return self.Procs.Len() == 0 && self.currProc == nil
}

func (self *CPU) Usage() float32 {
	return (float32(self.activeTimer) / float32(self.timer)) * 100.0
}

func (self *CPU) TurnaroundTime() float32 {
	totalTime := 0
	for _, x := range self.taTimes {
		totalTime += x
	}
	return (float32(totalTime) / float32(len(self.taTimes)))
}

func (self *CPU) WaitTime() float32 {
	totalTime := 0
	for _, x := range self.waitTimes {
		totalTime += x
	}
	return (float32(totalTime) / float32(len(self.waitTimes)))
}
