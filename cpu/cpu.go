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
	taTimes []int
	waitTimes []int
	arrivalTimes []int ///when we do priority scheduling, the arrival of a process is an event
}
//auxiliary function to check if sonething is an arrival time
func exists(l []int, key int) bool{
	for i := range l{ 
         if(l[i]==key){
			return true
		 }
	}
	return false
}
func NewCPU(schedAlgo sched.Scheduler) *CPU {
	procHeap := core.NewProcHeap(schedAlgo.Cmp)
	heap.Init(procHeap)

	return &CPU{
		timer:       0,//when we actually want round-robin, we set the quantum to a small number beforehand, for that only
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
func (self *CPU) changeProc() {
	fmt.Printf("Unloading proc PID: %d\n", self.currProc.Pid)
	copyProc := heap.Pop(self.procs).(*core.Proc)//the new guy, whom we'll run
	fmt.Printf("Loading proc PID: %d\n", copyProc.Pid)
	heap.Push(self.procs, *self.currProc)//push the old guy back
	self.currProc=copyProc
}
func (self *CPU) AddProc(proc core.Proc) {
	heap.Push(self.procs, proc)
	self.arrivalTimes=append(self.arrivalTimes, proc.Arrive)
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
		self.waitTimes=append(self.waitTimes, self.timer-self.currProc.Arrive)
	}
	if self.procDone(){

        self.taTimes=append(self.waitTimes, self.timer-self.currProc.Arrive)
		self.unloadProc()
	}

	if self.currProc != nil {
		self.currProc.Burst--
		self.activeTimer++
	}
}
//we need a different Tick() to implement preemption, since we can pop a process that isn't done, the above can't support it
func (self *CPU) PreemptiveTick() {
	self.timer++
    if self.available(){
       self.loadProc()
	}
	if self.procDone(){
		self.taTimes=append(self.waitTimes, self.timer-self.currProc.Arrive)
		self.waitTimes=append(self.waitTimes, self.timer-self.currProc.Arrive-self.currProc.InitBurst)
		self.unloadProc()
	}
	if exists(self.arrivalTimes, self.timer){ //I could do this in a much smarter way, by keeping track of indices, but let's keep it simple
		self.changeProc()
	}
	if self.currProc != nil{
		self.currProc.Burst--
		self.activeTimer++
	}
}
func (self *CPU) IsDone() bool {
	return self.procs.Len() == 0 && self.currProc == nil
}

func (self *CPU) Usage() float32 {
	return (float32(self.activeTimer) / float32(self.timer)) * 100.0
}

func (self *CPU) TurnaroundTime() float32{
    totalTime:=0
	for _, x := range(self.taTimes){
		totalTime+=x
	}
	return (float32(totalTime)/float32(len(self.taTimes)))
}

func (self *CPU) WaitTime() float32{
	totalTime:=0
	for _, x := range(self.waitTimes){
		totalTime+=x
	}
	return (float32(totalTime)/float32(len(self.waitTimes)))
}