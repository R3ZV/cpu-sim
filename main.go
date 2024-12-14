package main

import (
	"fmt"
	"sim/core"
	"sim/cpu"
	"sim/sched"
)

func addTestProcs(cpu *cpu.CPU) {
	cpu.AddProc(*core.NewProc(0, 5, 1))
	cpu.AddProc(*core.NewProc(3, 7, 10))
	cpu.AddProc(*core.NewProc(1, 9, 5))
	cpu.AddProc(*core.NewProc(2, 10, 3))
}

func main() {
	schedAlgs := []sched.Scheduler{sched.NewFCFS("FCFS"), sched.NewSJF("SJF")}

	for _, algo := range schedAlgs {
		cpu := cpu.NewCPU(algo)
		addTestProcs(cpu)

		fmt.Printf("Testing %s\n", algo.GetName())
		fmt.Println("============")
		fmt.Println()

		for !cpu.IsDone() {
			cpu.Tick()
		}

		fmt.Println()
		fmt.Println("====== CPU STATS ======")
		fmt.Printf("Usage: %.2f%%\n", cpu.Usage())
		fmt.Println()
	}
}
