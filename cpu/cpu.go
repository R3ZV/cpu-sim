package cpu

import (
	"sim/core"
	"sim/log"
	"sim/sched"

	"container/heap"
)

type CPU struct {
	timer          int
	activeTimer    int
	currProc       *core.Proc
	Procs          *core.ProcHeap
	eventFlag      bool //flag to alert us if we need to check for a switch
	PreemptiveFlag bool
	taTimes        []int
	waitTimes      []int
	arrivalTimes   []int
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
	procHeap := core.NewProcHeap(schedAlgo.Cmp, 0) //the CPU time always starts at 0
	heap.Init(procHeap)

	return &CPU{
		timer:       0,
		activeTimer: 0,
		currProc:    nil,
		Procs:       procHeap,
	}
}

func (self *CPU) AddJobs(jobs []core.Proc) {
	for _, job := range jobs {
		self.addProc(job)
	}
}

func (self *CPU) loadProc() {
	self.currProc = heap.Pop(self.Procs).(*core.Proc)
	log.Debug("[T=%d] Loading proc PID: %d\n", self.timer, self.currProc.Pid)
}

func (self *CPU) unloadProc() {
	log.Debug("[T=%d] Unloading proc PID: %d\n", self.timer, self.currProc.Pid)
	self.currProc = nil
}

func (self *CPU) changeProc() {
	log.Debug("[T=%d] Unloading proc PID: %d\n", self.timer, self.currProc.Pid)
	exchange := heap.Pop(self.Procs).(*core.Proc)
	log.Debug("[T=%d] Loading proc PID: %d\n", self.timer, exchange.Pid)
	heap.Push(self.Procs, *self.currProc)
	self.currProc = exchange
}

func (self *CPU) addProc(proc core.Proc) {
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
	}

	if self.procDone() {
		self.taTimes = append(self.taTimes, self.timer-self.currProc.Arrive)
		self.waitTimes = append(self.waitTimes, self.timer-self.currProc.Arrive-self.currProc.InitBurst)
		self.unloadProc()
		self.timer--
	}
	///there's no way to determine if it's an arrival otherwise
	if exists(self.arrivalTimes, self.timer) {
		self.eventFlag = true
	}
	if self.PreemptiveFlag == true && self.eventFlag == true {
		self.changeProc() //this might actually do nothing, since the process
		self.eventFlag = false
	}
	if self.currProc != nil {
		self.currProc.Burst--
		self.activeTimer++
	}
	self.timer++
	self.Procs.SetTime(self.timer)
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
		self.Procs.SetTime(self.timer)
	}

	// if exists(self.arrivalTimes, self.timer) {
	// 	self.changeProc()
	// }

	if self.currProc != nil {
		self.currProc.Burst--
		self.activeTimer++
	}
	self.timer++
	self.Procs.SetTime(self.timer)
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
		log.Debug("Around time: %d\n", x)
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
