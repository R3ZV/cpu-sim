package main

import (
	"fmt"
	"sim/core"
	"sim/cpu"
	"sim/log"
	"sim/sched"
)

func addJobs(workload *[][]core.Proc) {
	sched.FCFSJobs(workload)
	sched.SJFJobs(workload)
}

func main() {
	schedAlgs := []sched.Scheduler{
		sched.NewFCFS("FCFS"),
		sched.NewSJF("SJF"),
		sched.NewRMS("RMS"),
		// TODO:
		// sched.NewSJF("STCF"),
		// sched.NewSJF("RR"),
		// sched.NewPriority("Priority"),
		// sched.NewSJF("EDF"),
	}

	workload := [][]core.Proc{}
	addJobs(&workload)

	for _, algo := range schedAlgs {
		fmt.Printf("Testing %s\n", algo.GetName())
		fmt.Println("============")
		for i, jobs := range workload {
			fmt.Printf("Workload %d:\n", i)

			cpu := cpu.NewCPU(algo)
			cpu.AddJobs(jobs)

			for !cpu.IsDone() {
				if algo.IsPreemptive() {
					cpu.PreemptiveTick()
				} else {
					cpu.Tick()
				}
			}

			log.Assert(cpu.Procs.Len() == 0, "CPU hasn't finished its jobs")

			fmt.Printf("Usage: %.2f%%\n", cpu.Usage())
			fmt.Printf("Turnaround time: %.2f\n", cpu.TurnaroundTime())
			fmt.Printf("Waiting time: %.2f\n", cpu.WaitTime())
			fmt.Printf("\n")
		}

	}
}
