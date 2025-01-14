package cpu

import (
	"sim/core"
	"sim/log"
	"sim/sched"

	"container/heap"
)

type CPU struct {
	timer       int
	activeTimer int
	currProc    *core.Proc
	Procs       *core.ProcHeap

	// this event signifies that a new process has been added to the process queue
	// and in case of preemptiveness we should check if there is a more important
	// process
	event      bool
	preemptive bool

	// First map key is the Pid of a process, and the key of the second
	// map is one of "wait", "turnaround", "response" which represents
	// the compute time for each metric.
	//
	// In this case we assume that the Pids are unique and that there can't
	// be in the future another process with the Pid of a previous process.
	statTimes map[int]map[string]int
}

func NewCPU(schedAlgo sched.Scheduler) *CPU {
	procHeap := core.NewProcHeap(schedAlgo.Cmp, 0)
	heap.Init(procHeap)

	return &CPU{
		timer:       0,
		activeTimer: 0,
		currProc:    nil,
		Procs:       procHeap,
		event:       false,
		preemptive:  schedAlgo.IsPreemptive(),
		statTimes:   make(map[int]map[string]int),
	}
}

func (self *CPU) IsDone() bool {
	return self.Procs.Len() == 0 && self.currProc == nil
}

func (self *CPU) GetTimer() int {
	return self.timer
}

func (self *CPU) AddProc(proc core.Proc) {
	log.Assert(self.statTimes[proc.Pid] == nil, "There are multiple processes with same id at the same time")

	self.event = true
	heap.Push(self.Procs, proc)
	self.statTimes[proc.Pid] = make(map[string]int)
}

func (self *CPU) loadProc() {
	self.currProc = heap.Pop(self.Procs).(*core.Proc)
	self.statTimes[self.currProc.Pid]["response"] = self.timer - self.currProc.Arrive
	log.Debug("[T=%d] Loading proc PID: %d\n", self.timer, self.currProc.Pid)
}

func (self *CPU) unloadProc() {
	log.Debug("[T=%d] Unloading proc PID: %d\n", self.timer, self.currProc.Pid)
	self.currProc = nil
}

func (self *CPU) changeProc() {
	heap.Push(self.Procs, *self.currProc)
	optimalProc := heap.Pop(self.Procs).(*core.Proc)

	// context switch
	contextSwitchCost := 5
	if optimalProc != self.currProc {
		log.Debug("Context switch\n")
		log.Debug("[T=%d] Unloading proc PID: %d\n", self.timer, self.currProc.Pid)
		self.currProc = optimalProc
		self.timer += contextSwitchCost
		log.Debug("[T=%d] Loading proc PID: %d\n", self.timer, self.currProc.Pid)
	}
}

func (self *CPU) available() bool {
	return self.currProc == nil
}

func (self *CPU) procDone() bool {
	log.Assert(self.currProc != nil, "ProcDone should only be called if there was a process on CPU")
	return self.currProc.Burst == 0
}

func (self *CPU) Tick() {
	if self.preemptive && self.event {
		log.Debug("[T=%d] Preemptive: [%t] | Event: [%t]\n", self.timer, self.preemptive, self.event)
		if self.currProc != nil {
			self.changeProc()
		}
		self.event = false
	}

	if self.available() {
		self.loadProc()
	}

	if self.procDone() {
		self.statTimes[self.currProc.Pid]["turnaround"] = self.timer - self.currProc.Arrive
		self.statTimes[self.currProc.Pid]["wait"] = self.timer - self.currProc.Arrive - self.currProc.InitBurst

		self.unloadProc()
		self.timer--
	}

	if self.currProc != nil {
		self.currProc.Burst--
		self.activeTimer++
	}

	self.timer++
	self.Procs.SetTime(self.timer)
}

func (self *CPU) Usage() float32 {
	return (float32(self.activeTimer) / float32(self.timer)) * 100.0
}

func (self *CPU) TurnaroundTime() float32 {
	total := 0
	for _, stats := range self.statTimes {
		log.Debug("Around time: %d\n", stats["turnaround"])
		total += stats["turnaround"]
	}

	return (float32(total) / float32(len(self.statTimes)))
}

func (self *CPU) WaitTime() float32 {
	total := 0
	for _, stats := range self.statTimes {
		total += stats["wait"]
	}

	return (float32(total) / float32(len(self.statTimes)))
}

func (self *CPU) ResponseTime() float32 {
	total := 0
	for _, stats := range self.statTimes {
		total += stats["response"]
	}

	return (float32(total) / float32(len(self.statTimes)))
}
