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
func buildTestMetas(){
    ans=[]core.metaProc{}
	ans.append(core.newMetaProc(0, 5, 20))
	ans.append(core.newMetaProc(3, 7, 20))
	ans.append(core.newMetaProc(1, 9, 40))
	ans.append(core.newMetaProc(2, 10, 40))
	return ans
}
func main() {
	schedAlgs := []sched.Scheduler{sched.NewFCFS("FCFS"), sched.NewSJF("SJF"), sched.New}

	for _, algo := range schedAlgs {
		
		cpu := cpu.NewCPU(algo)
		addTestProcs(cpu)
        metas := buildTestMetas()
		fmt.Printf("Testing %s\n", algo.GetName())
		fmt.Println("============")
		fmt.Println()

		for !cpu.IsDone() {
            for i :=0; i<len(metas); i++{
				if(cpu.time % metas[i].Period ==0){
					cpu.AddProc(metas[i].Pid, metas[i].Burst, cpu.time)
				}
			}
			cpu.Tick()
		}

		fmt.Println()
		fmt.Println("====== CPU STATS ======")
		fmt.Printf("Usage: %.2f%%\n", cpu.Usage())
		fmt.Println()
	}
}
