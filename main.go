package main

import (
	"fmt"
	"sim/core"
	"sim/cpu"
	"sim/sched"
)

func addTestProcs(cpu *cpu.CPU) {
	cpu.AddProc(*core.NewProc(0, 5, 1, 1))
	cpu.AddProc(*core.NewProc(3, 7, 10, 1))
	cpu.AddProc(*core.NewProc(1, 9, 5, 1))
	cpu.AddProc(*core.NewProc(2, 10, 3, 2))
}

func main() {
	schedAlgs := []sched.Scheduler{sched.NewFCFS("FCFS"), sched.NewSJF("SJF"), sched.NewPriority("Priority")} //small issue I can't pin down with SRTF

	for _, algo := range schedAlgs {
		cpu := cpu.NewCPU(algo)
		addTestProcs(cpu)

		fmt.Printf("Testing %s\n", algo.GetName())
		fmt.Println("============")
		fmt.Println()
        if(algo.GetName() == "Priority" || algo.GetName() == "SRTF"){
			for !cpu.IsDone() {
                cpu.PreemptiveTick()
			}
		} else {
		   for !cpu.IsDone() {
			 cpu.Tick()
		   }
	    }
		fmt.Println("====== CPU STATS ======")
		fmt.Printf("Usage: %.2f%%\n", cpu.Usage())
		fmt.Printf("Turnaround time: %f\n", cpu.WaitTime()) 
        fmt.Printf("Waiting time: %f\n", cpu.TurnaroundTime()) 
		fmt.Println()
	}
}
