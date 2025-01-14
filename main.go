package main

import (
	"sim/core"
	"sim/cpu"
	"sim/log"
	"sim/math"
	"sim/sched"

	"fmt"
)

func addJobs(workload *[][]core.Proc) {
	sched.FCFSJobs(workload)
	sched.SJFJobs(workload)
	sched.PSJFJobs(workload)
	sched.PriorityJobs(workload)
}

func main() {
	schedAlgs := []sched.Scheduler{
		sched.NewFCFS("FCFS"),
		sched.NewSJF("SJF"),
		sched.NewPSJF("PSJF"),
		sched.NewPriority("Priority"),
		sched.NewRM("RM"),
		// sched.NewEDF("EDF"),
		// sched.NewSJF("RR"),
	}

	workload := [][]core.Proc{}
	addJobs(&workload)

	for _, algo := range schedAlgs {
		fmt.Printf("Testing %s\n", algo.GetName())
		fmt.Println("============")
		for i, jobs := range workload {
			duration := -1
			if algo.IsRealTime() {
				for _, job := range jobs {
					if job.Period == -1 {
						break
					}
					if duration == -1 {
						duration = 1
					}
					duration = math.LCM(duration, job.Period)
				}
			}

			log.Debug("LCM: %d\n", duration)

			fmt.Printf("Workload %d:\n", i)

			cpu := cpu.NewCPU(algo)
			procIdx := 0
			for !cpu.IsDone() || procIdx < len(jobs) {
				for procIdx < len(jobs) && jobs[procIdx].Arrive == cpu.GetTimer() {
					cpu.AddProc(jobs[procIdx])
					procIdx += 1
				}

				cpu.Tick()

				if duration != -1 && cpu.GetTimer() >= duration {
					break
				}
			}

			// log.Assert(cpu.Procs.Len() == 0, "CPU hasn't finished its jobs")

			fmt.Printf("Usage: %.2f%%\n", cpu.Usage())
			fmt.Printf("Turnaround time: %.2f\n", cpu.TurnaroundTime())
			fmt.Printf("Response time: %.2f\n", cpu.ResponseTime())
			fmt.Printf("Waiting time: %.2f\n", cpu.WaitTime())
			fmt.Printf("\n")
		}

	}
}
