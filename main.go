package main

import (
	"fmt"
	"sim/core"
	"sim/cpu"
	"sim/sched"
)

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func addJobs(workload *[][]core.Proc) {
	sched.FCFSJobs(workload)
	// SJFJobs(workload);
}

func loadJobs(cpu *cpu.CPU, jobs []core.Proc) {
	for _, job := range jobs {
		cpu.AddProc(job)
	}
}

func main() {
	schedAlgs := []sched.Scheduler{
		sched.NewFCFS("FCFS"),
		// sched.NewSJF("SJF"),
		// sched.NewPriority("Priority"),
	}

	workload := [][]core.Proc{}
	addJobs(&workload)

	for _, algo := range schedAlgs {
		fmt.Printf("Testing %s\n", algo.GetName())
		fmt.Println("============")
		for i, jobs := range workload {
			fmt.Printf("Workload %d:\n", i)

			cpu := cpu.NewCPU(algo)
			loadJobs(cpu, jobs)

			if algo.IsPreemptive() {
				for !cpu.IsDone() {
					cpu.PreemptiveTick()
				}
			} else {
				for !cpu.IsDone() {
					cpu.Tick()
				}
			}

			assert(cpu.Procs.Len() == 0, "CPU hasn't finished its jobs")

			fmt.Printf("Usage: %.2f%%\n", cpu.Usage())
			fmt.Printf("Turnaround time: %.2f\n", cpu.TurnaroundTime())
			fmt.Printf("Waiting time: %.2f\n", cpu.WaitTime())
			fmt.Printf("\n")
		}

	}
}
